package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	const method = "GET"
	const url = "https://eblog.fly.dev/index.html"
	var body io.Reader = nil
	req, err := http.NewRequestWithContext(context.TODO(), method, url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Set("User-Agent", "eblog/1.0")
	// wil be cannonicalized to Some-Key
	req.Header.Set("some-key", "a value")
	// wil overwrite the above since we usef Set rather than Add
	req.Header.Set("SOMe-KEY", "value")
	req.Write(os.Stdout)
}
