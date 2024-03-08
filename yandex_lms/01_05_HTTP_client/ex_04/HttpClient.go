package main

import (
  "io"
  "fmt"
  "net/http"
)

func SendRequest(url, role string, setRole bool) {
  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  if setRole {
    req.Header.Set("X-User-Role", role)
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("_%s\n", string(body))
}

func main() {
  SendRequest("http://127.0.0.1:8080/user", "user", false)
  SendRequest("http://127.0.0.1:8080/user", "user", true)
  SendRequest("http://127.0.0.1:8080/user", "noexist", true)

  SendRequest("http://127.0.0.1:8080/admin", "admin", false)
  SendRequest("http://127.0.0.1:8080/admin", "admin", true)
  SendRequest("http://127.0.0.1:8080/admin", "superadmin", true)
  SendRequest("http://127.0.0.1:8080/admin", "noexist", true)
}

