// download is a command-line tool to download a file from a URL.
// usage: download [-timeout duration] url filename

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
	"time"
)

func main() {
	dir := flag.String("dir", ".", "directory to save file")
	timeout := flag.Duration("timeout", 30*time.Second, "timeout for download")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("usage: download [-timeout duration] [-dir directory] url file")
	}
	url, filename := args[0], args[1]
	// always set a timeout when you make an HTTP request.
	c := http.Client{Timeout: *timeout}

	// dont't worry about the details of context for now;
	err := downloadAndSave(context.TODO(), &c, url, filepath.Join(*dir, filename))
	if err != nil {
		log.Fatal(err)
	}
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
