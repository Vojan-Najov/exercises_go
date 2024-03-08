package main

import "fmt"

func FindMaxKey(m map[int]int) int {
  iskey := false
  var max_key int
  var needed int
  for key, value := range m {
    if !iskey {
      iskey = true
      max_key = key
      needed = value
    } else if key > max_key {
      max_key = key
      needed = value
    }
  }
  return needed
}

func main() {
  m := map[int]int{1: 2, 2: 3, 7: 0, 8: 1}
  fmt.Println(FindMaxKey(m))
}
