package main

import (
  "fmt"
)

func Add(a, b float64) float64 {
  return a + b
}

func Multiply(a, b float64) float64 {
  return a * b
}

func PrintNumbersAscending(n int) {
  for i := 1; i <= n; i++ {
    fmt.Printf("%d", i)
    if i != n {
      fmt.Printf(" ")
    }
  }
  fmt.Println()
}

func main() {
  Add(1, 5)
  Multiply(4, 10)
  PrintNumbersAscending(10)
}

