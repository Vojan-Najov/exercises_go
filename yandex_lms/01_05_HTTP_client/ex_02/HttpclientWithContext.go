/*
HTTP Клиент с контекстом

Используя http.Client и context, напишите функцию
SendHTTPRequestWithContext(ctx context.Context, url string) (string, error)
которая делает GET-запрос к заданному URL и принимает контекст для
управления запросом.

Данная задача похожа на предыдущую, однако следует добавить контекст в
функцию, которая выполняет запрос. Ошибки обрабатывайте аналогично.

Для выполнения запроса используйте функцию NewRequestWithContext.

Примечания
Пример возвращаемой функцией строки:
"Hello, World!\n"
*/

package main

import (
  "io"
  "fmt"
  "time"
  "errors"
  "context"
  "net/http"
)

func SendHTTPRequestWithContext(ctx context.Context, url string) (string, error) {
  client := &http.Client{}

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return "", errors.New("Something went wrong...")
  }

  resp, err := client.Do(req.WithContext(ctx))
  if err != nil {
    return "", errors.New("Something went wrong...")
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return "", errors.New("Something went wrong...")
  }

  return string(body), nil
}

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  url := "noexist"
  str, err := SendHTTPRequestWithContext(ctx, url)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(str)
  }

  url = "https://ya.ru"
  str, err = SendHTTPRequestWithContext(ctx, url)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(str)
  }
}
