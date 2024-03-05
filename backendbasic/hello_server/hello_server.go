// The program demonstrates a minimal HTTP server that returns a 200 OK response
// with the text “hello, world”.

package main

import (
	"context"
	"net/http"
	"os"
)

func main() {
	server := http.Server{Addr: ":8080", Handler: TextHandler("hello, world!\r\n")}
	go server.ListenAndServe()

	req, _ := http.NewRequestWithContext(
		context.TODO(), "GET", "http://localhost:8080", nil,
	)
	resp, err := new(http.Client).Do(req)
	_ = err
	defer resp.Body.Close()

	resp.Write(os.Stdout)
}

// TextHandler is a simple http.Handler that returns a 200 OK response with
// the provided text.

type TextHandler string

var _ http.Handler = TextHandler("") // ensure TextHandler implements http.Handler

func (t TextHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(t)) // implicit 200 OK
}
