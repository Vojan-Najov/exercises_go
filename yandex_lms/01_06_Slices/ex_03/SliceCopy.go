/*
Копирование

Дан слайс целых чисел nums. Этот слайс имеет емкость больше его длины.
Создайте функцию SliceCopy(nums []int) []int, которая вернёт новый слайс
длиной и ёмкостью, равной длине nums. Скопируйте в него значения из
исходного слайса.

Примечания
Функцию main создавать не надо.
*/

package main

import "fmt"

func SliceCopy(nums []int) []int {
  copy := make([]int, len(nums), len(nums))
  for idx, el := range nums {
    copy[idx] = el
  }
  return copy
}

func main() {
  nums := make([]int, 10, 20)
  for i := 0; i < 10; i++ {
    nums[i] = i
  }
  copy := SliceCopy(nums)
  fmt.Println(len(nums), cap(nums), nums)
  fmt.Println(len(copy), cap(copy), copy)
}
