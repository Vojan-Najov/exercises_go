package main

import (
  "fmt"
)

func FindValue(nums []int) int {
  var result int
  m := map[int]bool{}

  for _, num := range nums {
    _, ok := m[num]
    if !ok {
      m[num] = false
    } else {
      m[num] = true
    }
  }

  for k, v := range m {
    if !v {
      result = k
      break
    }
  }

  return result
}

func main() {
  nums := []int {2, 2, 1}
  fmt.Println(FindValue(nums))
  nums = []int {2, 2, 1, 3, 1, 4, 3, 5, 4}
  fmt.Println(FindValue(nums))
}
