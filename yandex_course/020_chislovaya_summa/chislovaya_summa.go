package main

import (
  "fmt"
)

func SumDigitsRecursive(n int) int {
  if n >= 0 && n <= 9 {
    return n
  }
  if n >= -9 && n < 0 {
    return -n
  }
  digit := n % 10
  if (digit < 0) {
    digit *= -1
  }
  return digit + SumDigitsRecursive(n / 10)
}

func main() {
  for {
    var n int
    if _, err := fmt.Scanln(&n); err != nil {
      fmt.Println("Некорректный ввод")
      continue
    }
    fmt.Println(SumDigitsRecursive(n))
  }
}
