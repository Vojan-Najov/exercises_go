/*
Сколько времени прошло

Напишите функцию TimeAgo(pastTime time.Time) string, который вычисляет
время, прошедшее с момента заданного значения time.Time, и возвращает
удобочитаемую строку, указывающую, как давно это было.

Примечания

Пример работы функции:

    pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
    result := TimeAgo(pastTime)
    fmt.Println(result) // Output: 1 month ago

Пример работы с часами:

    pastTime := time.Now().Add(-2 * time.Hour)
    timeAgo := TimeAgo(pastTime)
    expected := "2 hours ago"
*/

package main

import (
  "fmt"
  "time"
)

func timeAgoAux(t int, s string) string {
  if t > 1 {
    s += "s"
  }
  return fmt.Sprintf("%d %s ago", t, s)
}
  

func TimeAgo(pastTime time.Time) string {
  diff := time.Now().Sub(pastTime)
  sec := int(diff.Seconds())
  years := sec / 60 / 60 / 24 / 365
  if years > 0 {
    return timeAgoAux(years, "year")
  }
  months := sec / 60 / 60 / 24 / 30
  if months > 0 {
    return timeAgoAux(months, "month")
  }
  days := sec / 60 / 60 / 24
  if days > 0 {
    return timeAgoAux(days, "day")
  }
  hours := sec / 60 / 60
  if hours > 0 {
    return timeAgoAux(hours, "hour")
  }
  minutes := sec / 60
  if minutes > 0 {
    return timeAgoAux(minutes, "minute")
  }
  return timeAgoAux(sec, "second")
}

func main() {
  pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
  result := TimeAgo(pastTime)
  fmt.Println(result) // Output: 1 month ago

  pastTime = time.Now().Add(-2 * time.Hour)
  timeAgo := TimeAgo(pastTime)
  fmt.Println(timeAgo)
}
