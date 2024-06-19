// Exercise 5.18: Without changing its behavior, rewrite the fetch function to use defer
// to close the writable file.

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		if _, _, err := fetch(url); err != nil {
			log.Printf("%s: %v", url, err)
		}
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		err = f.Close()
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
