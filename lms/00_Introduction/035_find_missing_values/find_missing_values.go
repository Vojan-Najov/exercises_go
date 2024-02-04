package main

import (
  "fmt"
)

func FindMissingValues(nums []int) []int {
  array := make([]bool, len(nums))

  for _, num := range nums {
    array[num - 1] = true
  }

  missing := []int{}
  for i, exist := range array {
    if !exist {
      missing = append(missing, i + 1)
    }
  }
  return missing
}

func main() {
  nums1 := []int {4, 3, 2, 7, 8, 2, 3, 1};
  fmt.Println(FindMissingValues(nums1));
  nums2 := []int {4, 4, 4, 4}
  fmt.Println(FindMissingValues(nums2));
  nums3 := []int {1, 1, 1}
  fmt.Println(FindMissingValues(nums3));
}
