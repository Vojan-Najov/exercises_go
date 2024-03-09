package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
  "errors"
  "strings"
  "time"

  "github.com/Vojan-Najov/exrcises_go/backendbasic/graduation/router"
  "github.com/Vojan-Najov/exrcises_go/backendbasic/graduation/ctxutil"
)

// register routes.
func buildBaseRouter() (*router.Router, error) {
  var r = new(router.Router)
  // --- design note: ---
  // you could just add the routes on a separate line for each.
  // but I like building the slice of routes and iterating over it;
  // it make the essential similarity of each route more obvious.
  for _, route := range []struct {
    pattern, method string
    handler         http.HandlerFunc
  }{
    // GET / returns "Hello, world!"
    {
      pattern: "/",
      method:  "GET",
      // ---- design note ----
      // this route demonstrates the simplest possible handler.
      // note the _ for the request parameter, we don't need it, so we dont't bind it.
      handler: func(w http.ResponseWriter, _ *http.Request) {
        w.Write([]byte("Hello, world!\r\n"))
      },
    },
    // GET /panic always panics.
    {
      pattern: "/panic",
      method:  "GET",
      // ---- design note ----
      // this route demonstrates how middleware can handle error conditions:
      // the panic will be caught by Recovery, which will write a 500 status
      // code and "internal server error" message to the response.
      // rather than leaving the connection hanging.
      // note the _ and _ for the request and response parametrs;
      // we don't need them, so we don't bind them.
      handler: func(_ http.ResponseWriter, _ *http.Request) {
        panic("oh my god JC, a bomb!")
      },
    },
    // POST /greet/json returns a JSON object with a greeting and
    // a category based on the age. It must be called with a JSON object in the form
    // {"first": "efron", "last": "licht", "age": 32}
    {
      pattern: "/greet/json",
      method:  "POST",
      // ---- design note ----
      // this route is a sophisticated example that has both path parameters
      // (using our custom router) and query parameters.
      // it 'puts everything together' and demonstrates how to use the router
      // and middleware together.
      handler: func(w http.ResponseWriter, r *http.Request) {
        var req struct {
          First, Last string
          Age         int
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
          // remember to return after writing an error!
          WriteError(w, err, http.StatusBadRequest)
          return
        }
        if req.Age < 0 {
          WriteError(w, errors.New("age must be >= 0"), http.StatusBadRequest)
          return
        }
        var category string
        switch {
        case req.Age < 13:
          WriteError(
            w, errors.New("forbidden: come back when you're older"),
            http.StatusForbidden,
          )
          return
        case req.Age < 21:
          category = "teenager"
        case req.Age > 65:
          category = "senior"
        default:
          category = "adult"
        }
        _ = WriteJSON(w, struct {
            Greeting string `json:"greeting"`
            Category string `json:"category"`
        }{
          fmt.Sprintf("Hello, %s %s!", req.First, req.Last),
          category,
        })
      },
    },
    // GET /time returns the current time in the given format.
    // it demonstrates how to use query parameters.
    {
      pattern: "/time",
      method:  "GET",
      handler: func(w http.ResponseWriter, r *http.Request) {
        format := r.URL.Query().Get("format")
        if format == "" {
          format = time.RFC3339
        }
        tz := r.URL.Query().Get("tz")
        var loc *time.Location = time.Local
        if tz != "" {
          var err error
          loc, err = time.LoadLocation(tz)
          if err != nil {
            WriteError(
              w, fmt.Errorf("invalid timezone %q: %w", tz, err),
              http.StatusBadRequest,
            )
            return
          }
        }
        _ = WriteJSON(w, struct {
          Time string `json:"time"`
        }{time.Now().In(loc).Format(format)})
      },
    },
    // Get/echo/{a}/{b}/{c} returns the path parameters as a JSON object in the form
    // {"a": "value of a", "b": "value of b", "c": "value of c"}
    // the query parameter "case" can be "upper" or "lower" to convert the values
    // uppercase or lowercase.
    {
      pattern: "/echo/{a:.+b}/{b:.+}/{c:.+}",
      method:  "GET",
      // ---- design note: ----
      // this route is sophisticated example that has both path parameters
      // (using our custom router) and query parameters.
      // It 'put everything together' and demonstrates how to use the router and
      // middleware together.
      handler: func(w http.ResponseWriter, r *http.Request) {
        vars, _ := ctxutil.Value[PathVars](r.Context())
        switch strings.ToLower(r.URL.Query().Get("case")) {
        case "upper":
          for k, v := range vars {
            vars[k] = strings.ToUpper(v)
          }
        case "lower":
          for k, v := range vars {
            vars[k] = strings.ToLower(v)
          }
        }
        _ = WriteJSON(w, vars)
      },
    },
  } {
      if err := r.AddRoute(route.pattern, route.handler, route.method); err != nil {
        return nil, fmt.Errorf(
          "AddRoute(%q, %v, %q) returned error: %v", route.pattern, route.handler,
          route.method, err,
        )
      }
      log.Printf("registered route: %s %s", route.method, route.pattern)
    }
  return r, nil
}
