package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

func GetGoogle(ctx context.Context) error {
	// deadline of the context is either 1 second from now, or the deadline
	// of the parent context, whichever is sooner.
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// always call cancel when you're done with the context to free associated resources
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if err != nil {
		return err
	}
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	resp.Write(os.Stdout)
	return nil
}

func main() {
	if err := GetGoogle(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
