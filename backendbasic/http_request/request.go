package main

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func main() {

	// The most basic example is a GET request with no body:
	{
		// use context.TODO() if you don't know what context to use
		ctx := context.TODO()

		// nil reader are OK; it means there's no body
		var body io.Reader = nil

		const method = "GET"

		const url = "http://eblog.fly.dev/index.html"

		// the function will parse the URL and set Host header;
		// invalid URLs will return an error.
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		_, _ = req, err
	}

	// For a POST request, you’ll need to provide a body.
	// The simplest way to do this is to use strings.NewReader to create a reader
	// from a string:
	{
		ctx := context.TODO()
		const method = "POST"
		const url = "https://eblog.fly.dev/index.html"
		var body io.Reader = strings.NewReader("hello, world")
		req, err := http.NewRequestWithContext(ctx, method, url, body)
		_, _ = req, err
	}
}
