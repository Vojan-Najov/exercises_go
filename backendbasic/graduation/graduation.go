package main

import (
  "flag"
  "log"
  "fmt"
  "net/http"

  "github.com/Vojan-Najov/exrcises_go/backendbasic/graduation/servermw"
)

func main() {
  port := flag.Int("port", 8080, "port to listen on")
  flag.Parse()

  h, err := buildBaseRouter()
  if err != nil {
    log.Fatal(err)
  }

  h = servermw.RecordResponse(h)
  h = servermw.Recovery(h)
  h = servermw.Log(h)
  h = servermw.Trace(h)

  // build and start the server
  server := http.Server{
    Addr:         fmt.Sprintf(":%d", *port),
    Handler:      h,
    ReadTimeout:  1 * time.Second,
    WriteTimeout: 1 * time.Second,
  }
  log.Printf("listening on %s", server.Addr)
  go server.ListenAndServe()
  time.Sleep(20 * time.Millisecond)
}
