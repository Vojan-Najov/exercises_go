package main

import (
  "fmt"
)

func main() {
  i := 2

  switch i {
  case 1:
    fmt.Println("i равно 1")
  case 2:
    fmt.Println("i равно 2")
  default:
    fmt.Println("i не равно ни 1, ни 2")
  }
}

