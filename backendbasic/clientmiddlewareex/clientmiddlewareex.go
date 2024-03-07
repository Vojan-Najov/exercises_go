// clientmiddlewareex makes a GET request to the specified URL, and prints
// the response body to stdout, using our middleware

package main

import (
  "context"
  "io"
  "log"
  "net/http"
  "os"
  "time"

  "github.com/Vojan-Najov/exercises_go/backendbasic/clientmiddlewareex/clientmw"
)

func main() {
  if len(os.Args) < 2 {
    log.Fatal("target url required")
  }
  target := os.Args[1]
  client := &http.Client{Transport: clientMiddleware(), Timeout: 5 * time.Second}
  req, err := http.NewRequestWithContext(context.TODO(), "GET", target, nil)
  if err != nil {
    log.Fatal(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  io.Copy(os.Stdout, resp.Body)
}

func clientMiddleware() http.RoundTripper {
  // specify the type as RoundTripFunc, not a http.RoundTripper, so thet we don't
  // have to repeatedly wrap it in RoundTripFunc(rt)
  var rt clientmw.RoundTripFunc
  const wait, tries = 10 * time.Millisecond, 3
  // first middleware applied will be the last one to run.
  // retry on 5xx status codes
  rt = clientmw.RetryOn5xx(http.DefaultTransport, wait, tries)
  // log request duration and status code; uses trace from next middleware
  rt = clientmw.Log(rt)
  // add trace id to request header
  rt = clientmw.Trace(rt)
  return rt
}
