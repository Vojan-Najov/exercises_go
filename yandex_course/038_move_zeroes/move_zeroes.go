package main

import (
  "fmt"
)

func MoveZeroes(nums []int) []int {
  n := len(nums)
  moved := make([]int, n)

  k := 0
  for _, num := range nums {
    if num != 0 {
      moved[k] = num
      k++
    }
  }

  return moved
}

func main() {
  nums := []int {0, 1, 0, 3, 12}
  fmt.Println(MoveZeroes(nums))
}
