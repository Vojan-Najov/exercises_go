/*
Слияние

Даны 2 слайса целых чисел nums1 и nums2. Создайте функцию
Join(nums1, nums2 []int) []int,
которя создаст новый слайс емкостью, вмещающей в себя ровно два слайса
(ёмкость должна быть равна его длине). Скопируйте в него сначала значения
nums1 затем nums2 и верните его.

Примечания
Функцию main создавать не надо.
*/

package main

import "fmt"

func Join(nums1, nums2 []int) []int {
  joined := make([]int, len(nums1) + len(nums2))
  var i int
  for _, el := range nums1 {
    joined[i] = el
    i++
  }
  for _, el := range nums2 {
    joined[i] = el
    i++
  }
  return joined
}

func main() {
  nums1 := []int{1, 2, 3}
  nums2 := []int{1, 2, 3, 4, 5}

  joined := Join(nums1, nums2)
  fmt.Println(len(joined), cap(joined), joined)
}
