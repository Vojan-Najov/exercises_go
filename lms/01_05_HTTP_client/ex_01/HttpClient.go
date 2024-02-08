/*
HTTP Клиент

Вам необходимо написать функцию
SendHTTPRequest(url string) (string, error),
которая делает GET-запрос к заданному URL и возвращает тело ответа в виде
строки.

Примечания
Нужный url подставляется внутри теста. В случае, если произошла ошибка,
необходимо возвращать ошибку: "Something went wrong..."

Пример ответа функции:
"Hello, World!\n"
*/

package main

import (
  "io"
  "fmt"
  "errors"
  "net/http"
)

func SendHTTPRequest(url string) (string, error) {
  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    return "", errors.New("Something went wrong...")
  }

  client := &http.Client{}
  resp, err := client.Do(req)
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
  url := "noexist"
  str, err := SendHTTPRequest(url)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(str)
  }

  url = "https://ya.ru"
  str, err = SendHTTPRequest(url)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(str)
  }


}
