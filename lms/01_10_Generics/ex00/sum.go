/*
Суммирование

Напишите универсальную функцию Sum[T Number](arr []T) T, которая может
суммировать элементы срезов разных числовых типов (например, int, float64).
Функция должна брать фрагмент любого числового типа и возвращать его сумму.

Примечания
Напишите так же констрейнт type Number interface, который будет обозначать
все численные типы.
*/

package main

import "fmt"

type Number interface {
  byte | int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[T Number](arr []T) T {
  var sum T
  for _, v := range arr {
    sum += v
  }
  return sum
}

func main() {
  fmt.Println(Sum([]int{1, 2, 3}))
  fmt.Println(Sum([]float64{1.1, 2.2, 3.3}))
}
