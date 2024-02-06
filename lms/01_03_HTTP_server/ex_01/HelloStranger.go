/*
Привет, странник

Напишите веб-сервер, который при обращении к нему будет возвращать
приветствие с именем пользователя, полученным из параметра запроса.

Если параметр пустой или отсутствует, сервер должен вернуть приветствие
hello stranger.

Если ответ содержит не только английские буквы, приветствие должно быть
hello dirty hacker.

Веб-сервер должен быть запущен на порту с номером 8080.

Формат ввода
Пример запроса:
curl localhost:8080/?name=John
# hello John
curl localhost:8080
# hello stranger

Примечания
Для проверки имени проще всего применить одну из функций пакета strings.
*/

package main

import (
  "fmt"
  "strings"
  "net/http"
)

type Handler struct {}

func IsNotAsciiLetter(r rune) bool {
  return !(r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z')
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Query().Get("name")
  if strings.ContainsFunc(name, IsNotAsciiLetter) {
    fmt.Fprintf(w, "hello %s\n", "dirty hacker")
  } else if len(name) != 0 {
    fmt.Fprintf(w, "hello %s\n", name)
  } else {
    fmt.Fprintf(w, "hello %s\n", "stranger")
  }
}

func main() {
  handler := &Handler{}
  if err := http.ListenAndServe(":8080", handler); err != nil {
    fmt.Println(err)
  }
}
