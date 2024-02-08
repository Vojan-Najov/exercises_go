/*
Middleware Conext

Вам необходимо создать веб-сервер с Middleware
RequestIDMiddleware(next http.Handler) http.Handler
для HTTP-обработчика
HelloHandler(w http.ResponseWriter, r *http.Request),
который будет добавлять информацию из заголовка "X-Request-ID" в контекст
запроса и затем использовать эту информацию в самом обработчике. Если
"X-Request-ID" не передается - необходимо написать об этом в формате в
виде "RequestID not found".

Не забудьте про функцию main, которая должна содержать методы для запуска
сервера в виде
func main() {
  http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))
  http.ListenAndServe(":8080", nil)
}

Примечания
Запрос:

    HTTP метод: GET
    Путь: "/hello"
    Заголовок: X-Request-ID: 12345
Ожидаемый ответ:
    Статус: 200 OK
    Тело ответа: "Hello! RequestID: 12345"

Запрос:
    HTTP метод: GET
    Путь: "/hello"
Ожидаемый ответ:
    Статус: 200 OK
    Тело ответа: "Hello! RequestID not found"
*/

package main

import (
  "fmt"
  "context"
  "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  requestID := r.Context().Value("RequestID")
  if requestID == nil {
    fmt.Fprintf(w, "Hello! RequestID not found\n")
  } else {
    fmt.Fprintf(w, "Hello! RequestID: %s\n", requestID.(string))
  }
}

func RequestIDMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    requestID := r.Header.Get("X-Request-ID")
    var ctx context.Context
    if len(requestID) == 0 {
      ctx = r.Context()
    } else {
      ctx = context.WithValue(r.Context(), "RequestID", requestID)
    }
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

func main() {
  http.Handle("/hello", RequestIDMiddleware(http.HandlerFunc(HelloHandler)))
  http.ListenAndServe(":8080", nil)
}
