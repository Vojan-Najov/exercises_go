package main

import (
  "fmt"
)

func Convert2D(nums []int, m, n int) [][]int {
  matrix := make([][]int, m)
  for i := 0; i < m; i++ {
    matrix[i] = make([]int, n)
    for j := 0; j < n; j++ {
      matrix[i][j] = nums[i * n + j]
    }
  }
  return matrix
}

func main() {
  nums := []int {1, 2, 3, 4, 5, 6, 7, 8}
  fmt.Println(Convert2D(nums, 4, 2))
}
