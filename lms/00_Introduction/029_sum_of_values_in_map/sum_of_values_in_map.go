package main

import "fmt"

func SumOfValuesInMap(m map[int]int) int {
  sum := 0
  for _, value := range m {
    sum += value
  }
  return sum
}

func main() {
  var m = map[int]int{1: 10, 2: 20, 3: 30}
  fmt.Println(SumOfValuesInMap(m))
}
