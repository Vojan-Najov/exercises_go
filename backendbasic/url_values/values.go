package main

import (
  "net/http"
  "context"
  "log"
  "os"
  "net/url"
)

func main() {
  const method = "Get"
  v := make(url.Values)
  // we use go's raw string syntax (`) to avoid having to escape the double quotes
  v.Add("q", `"of Emrakul"`)
  v.Add("order", "released")
  v.Add("dir", "asc")
  const path = "https://scryfall.com/search"
  // Encode() will escape the values for us. Remember the '?' separator.
  dst := path + "?" + v.Encode()
  req, err := http.NewRequestWithContext(context.TODO(), method, dst, nil)
  if err != nil {
    log.Fatal(err)
  }
  req.Write(os.Stdout)
}
