/*
Сервер Фибоначчи

Напишите веб-сервер, который будет запускаться на 8080 порту и принимать
запросы на получение следующего числа Фибоначчи, возвращая его значение.

Примеры запросов и ответов:
curl http://localhost:8080/
# 0
curl http://localhost:8080/
# 1
curl http://localhost:8080/
# 1
curl http://localhost:8080/
# 2

Сервер не сохраняет свое состояние между перезапусками. Таким образом,
если закрыть программу и запустить ее заново - подсчет начнется с 0.

Примечания

Числа Фибоначчи - это последовательность чисел, где каждое следующее число
равно сумме двух предыдущих чисел. Например, начиная с 0 и 1, первые
несколько чисел Фибоначчи будут выглядеть так: 0, 1, 1, 2, 3, 5, 8, 13 и
т.д. Эта последовательность часто встречается в математике, науке и
программировании. В вашей программе вы будете вычислять числа Фибоначчи и
возвращать их пользователю по запросу через веб-сервер.
*/

package main

import (
  "fmt"
  "net/http"
)

type Handler struct {
  fib1 int
  fib2 int
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%d\n", h.fib1)
  tmp := h.fib1
  h.fib1 = h.fib2
  h.fib2 += tmp
}

func main() {
  handler := &Handler{fib1: 0, fib2: 1}
  if err := http.ListenAndServe(":8080", handler); err != nil {
    fmt.Println(err)
  }
}