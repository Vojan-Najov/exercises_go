package main

import (
  "fmt"
  "strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
  var int1, int2 int
  var err error
  if int1, err = strconv.Atoi(a); err != nil {
    return 0, fmt.Errorf("invalid input, please provide two integers")
  }
  if int2, err = strconv.Atoi(b); err != nil {
    return 0, fmt.Errorf("invalid input, please provide two integers")
  }
  return int1 + int2, nil
}

func main() {
  if sum, err := SumTwoIntegers("12", "13"); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(sum)
  }
  if sum, err := SumTwoIntegers("12a", "13"); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(sum)
  }
  if sum, err := SumTwoIntegers("12", ""); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(sum)
  }
}

