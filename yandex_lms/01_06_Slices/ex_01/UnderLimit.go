/*
Дан неотсортированный слайс целых чисел. Напишите функцию
UnderLimit(nums []int, limit int, n int) ([]int, error),
которая будет возвращать первые n (либо меньше, если остальные не подходят)
элементов, которые меньше limit. В случае ошибки функция должна вернуть
nil и описание ошибки.

Примечания
Функцию main создавать не надо.
*/

package main

import (
  "fmt"
  "errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
  if nums == nil {
    return nil, errors.New("nums is nil")
  }
  if n < 0 {
    return nil, errors.New("Negative n")
  }

  under_limit := make([]int, 0, n)
  for i := 0; i < len(nums) && n > 0; i++ {
    if nums[i] < limit {
      under_limit = append(under_limit, nums[i])
      n--
    }
  }

  return under_limit, nil
}

func main() {
  nums := []int{1, 2, 3, 4, 1, 2, 3, 4}

  unlimit, _ := UnderLimit(nums, 10, 8)
  fmt.Println(unlimit)

  unlimit, _ = UnderLimit(nums, 10, 5)
  fmt.Println(unlimit)

  unlimit, _ = UnderLimit(nums, 3, 5)
  fmt.Println(unlimit)

  unlimit, _ = UnderLimit(nums, 3, 3)
  fmt.Println(unlimit)

  unlimit, _ = UnderLimit(nums, 3, 1)
  fmt.Println(unlimit)

  unlimit, _ = UnderLimit(nums, 0, 3)
  fmt.Println(unlimit)

  unlimit, _ = UnderLimit(nums, 2, 0)
  fmt.Println(unlimit)
}
