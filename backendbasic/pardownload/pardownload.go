// pardownload downloads a list of URLs in parallel, saving them to the specified dir.
// It exits with a nonzero status code if any of the downloads fail,
// where the status code is the number of failed downloads.

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	var dstDir string
	var client http.Client // the zero value of http.Client is a usable client.
	flag.StringVar(&dstDir, "dst", "", "destination directory; defaults to current dir")
	// set the timeout for the client using a command-line flag.
	flag.DurationVar(
		&client.Timeout, "timeout",
		1*time.Minute, "timeout for the request",
	)
	flag.Parse()

	src := flag.Args()
	if len(src) == 0 {
		log.Fatalf("can't copy")
	}
	// make the destination directory absolute, soour error messages are easier to read.
	dstDir, err := filepath.Abs(dstDir)
	if err != nil {
		log.Fatalf("invalid destination directory: %v", err)
	}
	// make a slice of the same length as src. so we can access it in parallel, without
	// worrying about synchronization.
	dst := make([]string, len(src))
	for i := range src {
		dst[i] = filepath.Join(dstDir, filepath.Base(src[i]))
	}

	errs := make([]error, len(src)) // si,ilarly, make a slice of errors

	// a WaitGroup waits for a collection of goroutines to finish.
	wg := new(sync.WaitGroup)
	wg.Add(len(src)) // add the number of goroutines we're going to wait for

	now := time.Now()
	for i := range src {
		i := i
		go func() {
			defer wg.Done() // tell the WaitGroup that we're done.

			// this is a simple function, so we dont't 'really' need to defer, but it's
			// a good habit to ger into.
			errs[i] = downloadAndSave(context.TODO(), &client, src[i], dst[i])
		}()
	}
	wg.Wait() // wait for all the goroutines to finish.

	log.Printf("downloaded %d files in %v", len(src), time.Since(now))
	var errCount int
	for i := range errs {
		if errs[i] != nil {
			log.Printf("err: %s -> %s: %v", src[i], dst[i], errs[i])
			errCount++
		} else {
			log.Printf("ok: %s -> %s", src[i], dst[i])
		}
	}
	os.Exit(errCount) // nonzero exit codes indicate failure.
}

func downloadAndSave(ctx context.Context, c *http.Client, url, dstPath string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("creating request: GET %q: %v", url, err)
	}
	// Do serialize a http.Request, sends it to the server,
	// and then deserializes the response to a http.Response.
	resp, err := c.Do(req)

	// always heck for errors after calling Do. errors from 'Do' usualy mean
	// something went wrong on the network
	if err != nil {
		return fmt.Errorf("request: %v", err)
	}
	defer resp.Body.Close() // always close response bodies when you're done with them.

	// immediately after checking for errors, check the response status code;
	// this is how the server tells us if the request succeded.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status: %s", resp.Status)
	}

	// ok, we have a succeful response. let's save it to a file

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("creating file: %v", err)
	}
	defer dstFile.Close() // always close files when you're done with them.
	if _, err := io.Copy(dstFile, resp.Body); err != nil {
		return fmt.Errorf("copying response to file: %v", err)
	}
	return nil
}
