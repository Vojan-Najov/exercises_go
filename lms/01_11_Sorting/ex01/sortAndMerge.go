/*
Сортировка со слиянием

Даны два слайса. Напишите программу, содержащую функцию
SortAndMerge(left, right []int) []int, которая объединит слайсы в один
отсортированный в два этапа: - отсортировать каждый слайс - объединить
полученные слайсы в один. Кстати, именно так работает алгоритм сортировки
слиянием ( merge sort)

Примечания

Ообъединять слайсы до сортировки не допустимо.
*/

package main

import (
  "slices"
  "fmt"
)

func SortAndMerge(left, right []int) []int {
  slices.Sort(left)
  slices.Sort(right)
  result := make([]int, len(left) + len(right))

  i := 0
  j := 0
  for i < len(left) && j < len(right) {
    if left[i] < right[j] {
      result[i + j] = left[i]
      i++
    } else {
      result[i + j] = right[j]
      j++
    }
  }
  for i < len(left) {
    result[i + j] = left[i]
    i++
  }
  for j < len(right) {
    result[i + j] = right[j]
    j++
  }
  return result
}

func main() {
  nums1 := []int{10, 4, 7}
  nums2 := []int{88, 5}

  nums := SortAndMerge(nums1, nums2)
  fmt.Println(nums)

  nums3 := []int{10, 4, 7}
  nums4 := []int{88, 5, 0, 99}

  nums = SortAndMerge(nums3, nums4)
  fmt.Println(nums)

  nums5 := []int{10, 4}
  nums6 := []int{88, 5}

  nums = SortAndMerge(nums5, nums6)
  fmt.Println(nums)
}
