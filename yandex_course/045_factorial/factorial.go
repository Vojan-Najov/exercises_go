package main

import (
  "fmt"
)

func Factorial(n int) (int, error) {
  if n < 0 {
    return 0, fmt.Errorf("factorial is not defined for negative numbers")
  } else if n < 2 {
    return 1, nil
  } else {
    fact := 1
    for n >= 2 {
      fact *= n
      n--
    }
    return fact, nil
  }
}

func main() {
  if f, err := Factorial(-1); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(f)
  }
  if f, err := Factorial(10); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(f)
  }
}

