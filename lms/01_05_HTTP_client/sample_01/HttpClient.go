package main

import (
  "io"
  "fmt"
  "net/http"
)

func main() {
  req, err := http.NewRequest(http.MethodGet, "https://ya.ru", nil)
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
