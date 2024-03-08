/*
Слайсы. Удаление элемента

Дан неотсортированный слайс целых чисел. Напишите функцию
Clean(nums []int, x int) ([]int),
которая удаляет из исходного слайса все вхождения x. Важно сохранить
порядок следования элементов и не использовать дополнительный слайс.

Примечания
Функцию main создавать не надо.
*/

package main

import "fmt"

func Clean(nums []int, x int) ([]int) {
  if nums == nil {
    return nil
  }

  size := len(nums)
  for i := 0; i < size; i++ {
    if nums[i] == x {
      for k, j := i, i + 1; j < len(nums); {
        nums[k] = nums[j]
        k++
        j++
      }
      i--
      size--
    }
  }

  return nums[:size]
}

func main() {
  nums := []int{1, 2, 3, 4, 3, 5, 2, 1}
  clean := Clean(nums, 4)
  fmt.Println(clean)

  nums = []int{1, 2, 3, 4, 3, 5, 2, 1}
  clean = Clean(nums, 2)
  fmt.Println(clean)

  nums = []int{1, 2, 3, 4, 3, 5, 2, 1}
  clean = Clean(nums, 3)
  fmt.Println(clean)

  nums = []int{1, 2, 3, 4, 3, 5, 2, 1}
  clean = Clean(nums, 1)
  fmt.Println(clean)

  nums = []int{1, 2, 3, 4, 3, 5, 2, 1}
  clean = Clean(nums, 5)
  fmt.Println(clean)

  nums = []int{1, 2, 3, 4, 3, 5, 2, 1}
  clean = Clean(nums, 6)
  fmt.Println(clean)

  nums = []int{5, 5}
  clean = Clean(nums, 5)
  fmt.Println(clean)

  nums = []int{3, -7}
  clean = Clean(nums, -7)
  fmt.Println(clean)
}
