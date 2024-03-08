/*
Обогащение ответа страннику *

Возьмите сервер, отвечающий hello stranger из этого урока.
Переделайте формат ответа на JSON вида:
{"greetings": "hello", "name": "stranger"}

При этом name берётся из запроса, а логика его подстановки не меняется.

Добавьте к этому middleware RPC(http.HandlerFunc), которая заменяет ответ
на формат:
{"status": "ok", "result": {"greetings": "hello", "name": "stranger"}}

Так же, переделайте Middleware Sanitize, чтобы она возвращала panic в
случае некорректного имени, и добавьте обработку этой паники в новой
middleware RPC так, чтобы отдавать пользователю ответ:
{"status": "error", "result": {}}

Примечания
Вот примеры запросов и соответствующих ответов:
Запрос: GET /hello?name=Alice Ответ:
{
    "greetings": "Hello",
    "name": "Alice"
}
Запрос: GET /hello?name="" Ответ: Ошибка 500, так как имя не соответствует
требованиям middleware Sanitize.
Ответ:
{
    "status": "error",
    "result": {}
}
*/

package main

import (
  "fmt"
  "strings"
  "encoding/json"
  "net/http"
)

var HelloHandler http.Handler = &HelloHandlerStruct{}
var data string
var status string

type Greeting struct {
  Greetings string `json:"greetings"`
  Name string      `json:"name"`
}

type HelloHandlerStruct struct {}

func (h * HelloHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  greet := Greeting{Greetings: "hello", Name: r.URL.Query().Get("name")}
  if d, err := json.Marshal(greet); err == nil {
    data = string(d)
  }
}

func SetDefaultName(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if len(name) == 0 {
      values := r.URL.Query()
      values.Set("name", "stranger")
      r.URL.RawQuery = values.Encode()
    }
    next.ServeHTTP(w, r)
  })
}

func Sanitize(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if strings.ContainsAny(name, "!@#$%^&*()_+") {
      panic("error")
    } else {
      next.ServeHTTP(w, r)
    }
  })
}

func RPC(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    defer func() {
      if err := recover(); err != nil {
        fmt.Fprintf(w, "{\"status\":\"error\",\"result\":{}}")
      } else {
        fmt.Fprintf(w, "{\"status\":\"ok\",\"result\":")
        fmt.Fprintf(w, "%s", data)
        fmt.Fprintf(w, "}")
      }
    }()
    next.ServeHTTP(w, r)
  })
}

func main() {
  mux := http.NewServeMux()
  mux.Handle("/hello", RPC(SetDefaultName(Sanitize(HelloHandler))))
  if err := http.ListenAndServe(":8080", mux); err != nil {
    fmt.Println(err)
  }
}

