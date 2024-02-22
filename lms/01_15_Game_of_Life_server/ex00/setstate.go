/*

Задание состояния

В этом уроке мы заполняем сетку живыми клетками при старте программы.
Давайте сделаем отдельный адрес, при обращении по которому будет задано
новое состояние игры.
Напишите HTTP server, принимающий запросы по адресу
http://localhost:8081/setstate (метод POST) с новым процентом заполнения
сетки. Процент заполнения должен быть передан в формате JSON:
{fill: 30} - это означает заполнить сетку на 30%. Заданный процент
заполнения сохраняйте в файл state.cfg в формате: 30%.

Примечания

Пример запроса:
curl -X POST http://localhost:8081/setstate
-H 'Content-Type: application/json'
-d '{"fill":30}'

*/

package main

import (
  "io"
  "os"
  "fmt"
  "slices"
  "net/http"
  "encoding/json"
)

type Config struct {
  Fill int `json:"fill"`
}

func SetState(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" && slices.Contains(r.Header["Content-Type"], "application/json") {
    data, err  := io.ReadAll(r.Body)
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      return 
    }
    var config Config
    err = json.Unmarshal(data, &config)
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      return
    }
    f, err := os.OpenFile("state.cfg",
                          os.O_CREATE | os.O_TRUNC | os.O_WRONLY,
                          0644)
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      return
    }
    defer f.Close()
    fmt.Fprintf(f, "%d%%", config.Fill)
  }
}

func main() {
  mux := http.NewServeMux()

  mux.Handle("/setstate", http.HandlerFunc(SetState))
  if err := http.ListenAndServe(":8081", mux); err != nil {
    fmt.Fprintln(os.Stderr, err)
  }
}
