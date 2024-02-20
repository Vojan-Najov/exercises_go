/*
Следующий рабочий день

Напишите функцию NextWorkday(start time.Time) time.Time, которая вычисляет
дату следующего рабочего дня (исключая выходные). Учитывая дату начала,
функция должна возвращать дату следующего рабочего дня.

Примечания
Рабочая неделя начинается с понедельника :)

*/

package main

import (
  "fmt"
  "time"
)

func NextWorkday(start time.Time) time.Time {
  workday := start.Add(24 * time.Hour)
  if workday.Weekday() == time.Saturday {
    workday = workday.Add(24 * time.Hour)
  }
  if workday.Weekday() == time.Sunday {
    workday = workday.Add(24 * time.Hour)
  }
  return workday
}

func main() {
  fmt.Println(time.Now())
  fmt.Println(NextWorkday(time.Now()))
  fmt.Println(NextWorkday(time.Now().Add(-24 * time.Hour)))
  fmt.Println(NextWorkday(time.Now().Add(-48 * time.Hour)))
  fmt.Println(NextWorkday(time.Now().Add(-72 * time.Hour)))
  fmt.Println(NextWorkday(time.Now().Add(-96 * time.Hour)))
}

