/*
Восстановление состояния

Добавьте возможность восстановления последнего заданного состояния сетки.
Напишите HTTP server, принимающий запросы по адресу
http://localhost:8081/reset (метод PUT).
При обращении по данному адресу сервер должен прочитать файл state.cfg
(формат в предыдущем задании) и вернуть пользователю последний сохраненный
процент заполнения.

Примечания

Пример ответа: {"fill":30}

*/

package main

import (
  "os"
  "fmt"
  "net/http"
  "encoding/json"
)

type Config struct {
  Fill int `json:"fill"`
}

func ResetState(w http.ResponseWriter, r *http.Request) {
  if r.Method != "PUT" {
    return
  }

  data, err := os.ReadFile("state.cfg")
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }

  var conf Config
  _, err = fmt.Sscanf(string(data), "%d%%", &conf.Fill)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }

  data, err = json.Marshal(conf)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }

  fmt.Fprint(w, string(data)) 
}

func main() {
  mux := http.NewServeMux()

  mux.Handle("/reset", http.HandlerFunc(ResetState))

  if err := http.ListenAndServe(":8081", mux); err != nil {
    fmt.Fprintln(os.Stderr, err)
  }
}
