package main

import "fmt"

func main() {
  var meters float64 = 0.0
  fmt.Scan(&meters)

  const exchangeRate float64 = 1852.0
  seaMiles := meters / exchangeRate
  fmt.Println(seaMiles)
}
