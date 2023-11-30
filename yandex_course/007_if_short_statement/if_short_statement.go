package main

import (
  "fmt"
)

func main() {
  number := 5
  if square := number * number; square > 50 {
    fmt.Println("Квадрат числа больше 50")
  } else {
    fmt.Println("Квадрат числа меньше или равен 50")
  }
}
