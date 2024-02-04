package main

import "fmt"

func ContainsDuplicate(nums []int) bool {
  set := make(map[int]bool)
  for _, num := range nums {
    if _, ok := set[num]; ok {
      return true
    }
    set[num] = true
  }
  return false
}

func main() {
  nums1 := []int{1, 2, 3, 4, 5}
  fmt.Println(ContainsDuplicate(nums1));

  nums2 := []int{1, 2, 3, 4, 5, 1}
  fmt.Println(ContainsDuplicate(nums2));
}
