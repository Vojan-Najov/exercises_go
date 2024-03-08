/*
Сервер Фибоначчи с метриками

Напишите веб-сервер, который будет считать метрики времени ответа сервиса.

Возьмите в качестве основы веб-сервер из предыдущей задачи, вычисляющий
числа Фибоначчи, и добавьте к нему хендлер /metrics, который отдаёт:

rpc_count 10 где 10 — число раз, которое вызвали хендлер,
отвечающий числами Фибоначчи. Выглядит это так:

curl http://localhost:8080/metrics
# rpc_count 0
curl http://localhost:8080/
# 0
curl http://localhost:8080/metrics
# rpc_count 1
# ... 
curl http://localhost:8080/
# 34
curl http://localhost:8080/metrics
# rpc_count 10
*/

package main

import (
  "fmt"
  "net/http"
)

type Handler struct {
  fib1 int
  fib2 int
  rpc_count int
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  if path == "/" {
    fmt.Fprintf(w, "%d\n", h.fib1)
    tmp := h.fib1
    h.fib1 = h.fib2
    h.fib2 += tmp
    h.rpc_count++
  } else if path == "/metrics" {
    fmt.Fprintf(w, "rpc_count %d\n", h.rpc_count)
  }
}

func main() {
  handler := &Handler{fib1: 0, fib2: 1, rpc_count: 0}
  if err := http.ListenAndServe(":8080", handler); err != nil {
    fmt.Println(err)
  }
}
