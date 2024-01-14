package main

import (
  "fmt"
  "strings"
)

func binary_search(sl []string, prefix string) (string, bool) {
  low := 0
  high := len(sl) - 1
  for low <= high {
    mid := (low + high) / 2
    midval := sl[mid]
    if strings.HasPrefix(midval, prefix) {
      for mid--; mid >= 0; mid-- {
        midval = sl[mid]
        if !strings.HasPrefix(midval, prefix) {
          mid++
          return sl[mid], true
        }
      }
      return midval, true
    } else if midval < prefix {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }
  return "Не найдено", false
}

func main() {
  var n int
  if _, err := fmt.Scanln(&n); err != nil || n < 0 {
    fmt.Println("Incorrect input")
    return
  }

  book := make([]string, n)
  for i := 0; i < n; i++ {
    var s1, s2 string
    fmt.Scanln(&s1, &s2)
    book[i] = s1 + " " + s2
  }

  for {
    var s1, s2 string
    fmt.Scanln(&s1, &s2)
    if (len(s1) == 0) {
      break
    }
    if (len(s2) > 0) {
      s1 = s1 + " " + s2
    }
    s, _ := binary_search(book, s1)
    fmt.Println(s)
  }
}
