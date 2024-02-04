package main

import (
  "fmt"
)

func main() {
  var n, m int
  if _, err := fmt.Scanln(&n); err != nil || n < 0 {
    fmt.Println("Incorrect input")
    return
  }
  if _, err := fmt.Scanln(&m); err != nil || m < 0 {
    fmt.Println("Incorrect input")
    return
  }

  if n == 0 {
    m++;
  }
  for ; n > 0; n-- {
    for i := n; i > 0; i-- {
      fmt.Print("*")
    }
    fmt.Println()
  }
  for m--; m > 0; m-- {
    fmt.Println("*")
  }
}
