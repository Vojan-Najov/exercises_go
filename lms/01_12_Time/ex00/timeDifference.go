/*
Разница времени

Напишите функцию TimeDifference(start, end time.Time) time.Duration,
которая вычисляет разницу во времени между двумя заданными моментами
времени. Функция должна принимать на вход два значения time.Time и
возвращать продолжительность между ними.
Примечания

Пример работы функции
    start := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
    end := time.Date(2023, 10, 23, 4, 41, 49, 0, time.UTC)
    diff := TimeDifference(start, end)
    fmt.Println(diff) // Output: 2h0m0s
*/

package main

import (
  "fmt"
  "time"
)

func TimeDifference(start, end time.Time) time.Duration {
  return end.Sub(start)
}

func main() {
  start := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
  end := time.Date(2023, 10, 23, 4, 41, 49, 0, time.UTC)
  diff := TimeDifference(start, end)
  fmt.Println(diff) // Output: 2h0m0s
}
