package main

import "fmt"

func main() {
  var age int
  fmt.Println("Введите ваш возраст:")
  _, err := fmt.Scanln(&age)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(age)
}
