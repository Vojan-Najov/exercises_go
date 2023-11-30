package main

import "fmt"

func main() {
  var num int

  _, err := fmt.Scanln(&num)
  if err != nil {
    fmt.Println(err);
  }

  doubled := num * 2
  fmt.Println("Удвоенное число:", doubled)
}
