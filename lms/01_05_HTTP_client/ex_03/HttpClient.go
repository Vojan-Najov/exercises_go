package main

import (
  "io"
  "fmt"
  "net/http"
)

func main() {
  {
    req, err := http.NewRequest(http.MethodGet,
                                "http://127.0.0.1:8080/hello", nil)
    if err != nil {
     fmt.Println(err.Error())
     return
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
     fmt.Println(err.Error())
     return
    }
    // читаем тело ответа
    body, err := io.ReadAll(resp.Body)
    if err != nil {
     fmt.Println(err)
     return
    }

    // выводим тело ответа на экран
    fmt.Println(string(body))
  }
  {
    req, err := http.NewRequest(http.MethodGet,
                                "http://127.0.0.1:8080/hello", nil)
    if err != nil {
     fmt.Println(err.Error())
     return
    }
    req.Header.Set("X-Request-ID", "123")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
     fmt.Println(err.Error())
     return
    }
    // читаем тело ответа
    body, err := io.ReadAll(resp.Body)
    if err != nil {
     fmt.Println(err)
     return
    }

    // выводим тело ответа на экран
    fmt.Println(string(body))
  }
}
