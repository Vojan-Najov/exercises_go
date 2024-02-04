package main

import (
  "fmt"
)

func IntToBinary(num int) (string, error) {
  if num < 0 {
    return "", fmt.Errorf("negative numbers are not allowed")
  } else if num == 0 {
    return "0", nil
  }

  str := ""
  for num != 0 {
    if num & 1 == 1 {
      str = "1" + str
    } else {
      str = "0" + str
    }
    num >>= 1
  }
  return str, nil
}

func main() {
  if str, err := IntToBinary(8); err == nil {
    fmt.Println(str)
  } else {
    fmt.Println(err)
  }
  if str, err := IntToBinary(7); err == nil {
    fmt.Println(str)
  } else {
    fmt.Println(err)
  }
  if str, err := IntToBinary(0); err == nil {
    fmt.Println(str)
  } else {
    fmt.Println(err)
  }
  if str, err := IntToBinary(-2); err == nil {
    fmt.Println(str)
  } else {
    fmt.Println(err)
  }
}
