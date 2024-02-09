/*
Слияние двух частей

Дан слайс nums, состоящий из 2n элементов в формате
[x0,x1,...,xn,y0,y1,...,yn].
Создайте функцию Mix(nums []int) []int, которая вернёт слайс, содержащий
значения в следующем порядке: [x0,y0,x1,y1,...,xn,yn].

Примечания
Функцию main создавать не надо.
*/

package main

import "fmt"

func Mix(nums []int) []int {
  i, j, k := 0, len(nums) / 2, 0
  if len(nums) % 2 != 0 {
    j++
  }

  mix := make([]int, len(nums))
  for j < len(nums) {
    mix[k] = nums[i]
    k++
    i++
    mix[k] = nums[j]
    k++
    j++
  }
  if len(nums) % 2 != 0 {
    mix[k] = nums[len(nums) / 2]
  }

  return mix
}

func main() {
  nums := []int{1, 2, 3, 4, 10, 20, 30, 40}
  mix := Mix(nums)
  fmt.Println(nums)
  fmt.Println(mix)

  nums = []int{1, 2, 3, 4, 5}
  mix = Mix(nums)
  fmt.Println(nums)
  fmt.Println(mix)
}
