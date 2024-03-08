package main

import (
  "fmt"
)

func main() {
  var m map[int]bool
  m = make(map[int]bool)
  var n int

  if _, err := fmt.Scanln(&n); err != nil || n < 0 {
    fmt.Println("Incorrect input")
    return
  }

  for i := 1; i <= n; i++ {
    is_prime := true
    for j := 2; j * j <= i; j++ {
      if i % j == 0 {
        is_prime = false
        break
      }
    }
    m[i] = is_prime
  }

  for i:= 3; i <= n; i = i + 5 {
    if m[i] {
      fmt.Print("хоп")
    } else {
      fmt.Print(i)
    }
    if (i + 4 <= n) {
      fmt.Print(" ")
    }
  }

  fmt.Println()
}
