package clientmw

import (
  "os"
  "log"
  "net/http"
  "time"
  "errors"
  "syscall"
  "fmt"

  "github.com/google/uuid"
  "github.com/Vojan-Najov/exercises_go/backendbasic/clientmiddlewareex/ctxutil"
  "github.com/Vojan-Najov/exercises_go/backendbasic/clientmiddlewareex/trace"
)

// RoundTripFunc is an adapter to allow the use of ordinary functions as 
// RoundTrippers, a-la http.HandlerFunc
type RoundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip implements the RoundTripper interface by calling f(r)
func (f RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
  return f(r)
}

// assert that RoundTripFunc implements http.RoundTripper at compile time
var _ http.RoundTripper = RoundTripFunc(nil)

// Client Middleware without Dependencies

// we'll use this helper function to log the beginning and end of each middleware.
// no need for this in the real world, but it should help you understand what's
// going on
func logExec(name string) func() {
  log.Printf("middleware: begin %s", name)
  return func() {
    defer log.Printf("middleware: end %s", name)
  }
}

// TimeRequest returns a RoundTripFunc that logs the durations of the request.
func TimeRequst(rt http.RoundTripper) RoundTripFunc {
  return func(r *http.Request) (*http.Response, error) {
    // for demonstration purposes, we'll add these logs to each middleware;
    // don't do this in production!
    defer logExec("TimeRequest")()

    start := time.Now()
    // call next middleware, or http.DefaultTransport.RoundTrip if this is the last
    // middleware
    resp, err := rt.RoundTrip(r)
    if err != nil {
      log.Printf("%s %s: errored after %s", r.Method, r.URL, time.Since(start))
      return nil, err
    }
    log.Printf(
      "%s %s: %d %s in %s", r.Method, r.URL,
      resp.StatusCode, http.StatusText(resp.StatusCode), time.Since(start),
    )
    return resp, nil
  }
}

// Client Middleware with Injected Dependencies

// RetryOn5xx returns a RoundTripFunc that retries the request up to n times
// if the server returns a 5xx status code.
// It will use exponential backoff: first retry will be after wait,
// second after 2 * wait, third after 4 * wait, etc.
func RetryOn5xx(rt http.RoundTripper, wait time.Duration, tries int) RoundTripFunc {
  // validate arguments OUTSIDE of the closure, so that it only happens once
  if tries <= 1 {
    panic("tries must be > 1")
  }
  if wait <= 0 {
    panic("wait must be > 0")
  }
  return func(r *http.Request) (*http.Response, error) {
    defer logExec("RetryOn5xx")()
    //retry logic
    var retryErrs error
    for retry := uint(0); retry < uint(tries); retry++ {
      if retry > 0 {
        time.Sleep(wait << retry)
      }
      // call next middleware, or http.DefaultTransport.RondTrip if this is the last
      // middleware
      resp, err := rt.RoundTrip(r)
      if errors.Is(err, syscall.ECONNREFUSED) ||
         errors.Is(err, syscall.ECONNRESET) {
        retryErrs = errors.Join(retryErrs, err)
        continue
      }
      if err != nil {
        return nil, fmt.Errorf("failed after %d retries: %v", retry, err)
      }
      switch sc := resp.StatusCode; {
      case sc >= 200 && sc < 400: // success! we're done here
        return resp, nil
      case sc <= 400 && sc < 500: // 4xx status code
        return nil, fmt.Errorf("failed after %d retries: %s", retry, resp.Status)
      default:  // 5xx, 1xx, or unknown status code
        retryErrs = errors.Join(
          retryErrs, fmt.Errorf("try %d: %s", retry, resp.Status),
        )
      }
    }
    return nil, fmt.Errorf("failed after %d retries: %v", tries, retryErrs)
  }
}

// Trace returns a RoundTripFunc that
// - adds a trace to th request context
// - generating a new one if necessary
// - adds the X-Trace-ID and X-Request-ID headers to the request
// - then calls the next RoundTripper
func Trace(rt http.RoundTripper) RoundTripFunc {
  return func(r *http.Request) (*http.Response, error) {
    defer logExec("Trace")()
    // does the request already have a trace? if so, use it. otherwise, generate a new
    traceID, err := uuid.Parse(r.Header.Get("X-Trace-ID"))
    if err != nil {
      traceID = uuid.New()
    }

    // build the trace. it's a small struct, so we put it directly in the context and
    // don't bother with a pointer.
    trace := trace.Trace{TraceID: traceID, RequestID: uuid.New()}

    // add trace to context; retrive with ctxutil.Value[Trace](ctx)
    ctx := ctxutil.WithValue(r.Context(), trace)
    // add context to request
    r = r.WithContext(ctx)

    // add trace id & request id to headers
    r.Header.Set("X-Trace-ID", trace.TraceID.String())
    r.Header.Set("X-Request-ID", trace.RequestID.String())
    // call next middleware, or http.DefaultTransport.RoundTrip if this is the last
    return rt.RoundTrip(r)
  }
}

// Log returns a RoundTripFunc that logs the request duration and status code.
// It uses the trace from the context as a prefix, if it exists.
func Log(rt http.RoundTripper) RoundTripFunc {
  return func(r *http.Request) (*http.Response, error) {
    defer logExec("Log")()
    trace, ok := ctxutil.Value[trace.Trace](r.Context())
    var prefix string
    if ok {
      prefix = fmt.Sprintf(
        "%s %s: [%s %s]: ",
        r.Method, r.URL, trace.TraceID, trace.RequestID,
      ) 
    } else {
      prefix = fmt.Sprintf("%s %s: ", r.Method, r.URL)
    }

    logger := log.New(os.Stderr, prefix, log.LstdFlags | log.Lshortfile)
    // add logger to context; retrieve with ctxutil.Value[log.Logger](ctx)
    ctx := ctxutil.WithValue[*log.Logger](r.Context(), logger)
    r = r.WithContext(ctx)

    start := time.Now()
    // call next middleware, or http.DefaultTransport.RoundTrip if this is the last
    resp, err := rt.RoundTrip(r)
    if err != nil {
      logger.Printf("errored after %s: %s", time.Since(start), err)
      return nil, err
    }
    logger.Printf(
      "%d %s in %s",
      resp.StatusCode, http.StatusText(resp.StatusCode), time.Since(start),
    )
    return resp, nil
  }
}
