package main

import (
  "fmt"
)

func DivideIntegers(a, b int) (float64, error) {
  if b == 0 {
    return 0.0, fmt.Errorf("division by zero is not allowed")
  }
  return float64(a)/float64(b), nil
}

func main() {
  if n, err := DivideIntegers(10, 0); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Result: ", n)
  }
  if n, err := DivideIntegers(10, 3); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Result: ", n)
  }
}
