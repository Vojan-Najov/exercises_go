package main

import (
  "context"
  "fmt"
  "net/http"
  "time"
)

func main() {
  // создаём контекст с таймаутом в 5 секунд
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  // создаём клиент
  client := &http.Client{}

  // создаём запрос
  req, err := http.NewRequest("GET", "http://ya.ru", nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  // отправляем запрос с контекстом
  resp, err := client.Do(req.WithContext(ctx))
  if err != nil {
    fmt.Println(err)
    return
  }
  // в этом примере мы не собираемся читать тело ответа
  defer resp.Body.Close()

  // обрабатываем ответ
  fmt.Println(resp.StatusCode)
}
