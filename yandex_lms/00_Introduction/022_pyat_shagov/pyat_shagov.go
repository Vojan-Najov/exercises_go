package main

import "fmt"

func ReverseSlice(slice []int) []int {
  reverse_slice := make([]int, len(slice))

  j := 0;
  for i := len(slice) - 1; i >= 0; i-- {
    reverse_slice[j] = slice[i]
    j++
  }
  return reverse_slice
}

func main() {
  slice := make([]int, 19)
  for i := 0; i < 19; i++ {
    slice[i] = i
  }
  fmt.Println(slice)
  reverse_slice := ReverseSlice(slice)
  fmt.Println(reverse_slice)
}
