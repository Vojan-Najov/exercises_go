/*
Сортировка слайса uint

Напишите программу, содержащую функцию SortNums(nums []uint),
которая сортирует слайс nums по возрастанию
*/

package main

import (
  "fmt"
  "slices"
)

func SortNums(nums []uint) {
  slices.Sort(nums)
}

func main() {
  nums := []uint{10, 1, 8, 4, 3, 2}
  SortNums(nums)
  fmt.Println(nums)
}
