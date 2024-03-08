/*
Значение выражения

Функция MaxExpressionValue(nums []int) int принимает на вход слайс nums.

Найдите максимальное значение выражения
nums[s] — nums[r] + nums[q] — nums[p], где p, q, r и s — индексы слайса,
a s > r > q > p

Например, для nums := []int{3, 9, 10, 1, 30, 40} функция должна вернуть
значение 46 (поскольку 40 – 1 + 10 – 3 - максимально).

Задачу надо решить, используя принципы динамического программирования.
Примечания
В качестве решения надо отправить функцию MaxExpressionValue и все
вспомогательные функции, которые вам потребуются.
*/

package main

import "fmt"

func MaxExpressionValue(nums []int) int {
  first := make([]int, len(nums) + 1)
  for i := len(nums) - 1; i >= 0; i-- {
    first[i] = max(first[i + 1], nums[i])
  }

  fmt.Println(first)

  second := make([]int, len(nums))
  for i := len(nums) - 2; i >= 0; i-- {
    second[i] = max(second[i + 1], first[i + 1] - nums[i]) 
  }

  fmt.Println(second)

  third := make([]int, len(nums) - 1)
  for i := len(nums) - 3; i >= 0; i-- {
    third[i] = max(third[i + 1], second[i + 1] + nums[i])
  }
  fmt.Println(third)

  fourth := make([]int, len(nums) - 2)
  for i := len(nums) - 4; i >= 0; i-- {
    fourth[i] = max(fourth[i + 1], third[i + 1] - nums[i]) 
  }

  fmt.Println(fourth)

  return fourth[0]
}

func main() {
  nums := []int{3, 9, 10, 1, 30, 40}
  MaxExpressionValue(nums)
}
