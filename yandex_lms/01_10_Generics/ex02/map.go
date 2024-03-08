/*
Map

Создайте дженерик функцию
Map[T, U any](arr []T, transform func(T) U) []U,
которая применяет заданную функцию преобразования к каждому элементу среза
и возвращает новый срез с преобразованными значениями. Функция должна
работать со срезами любого типа.

Примечания
Пример работы функции

arr := []int{1, 2, 3}
result := Map(arr, func(x int) string {
    return fmt.Sprintf("%d!", x)
})
fmt.Println(result) // Output: [1! 2! 3!]
*/
 
package main

import "fmt"

func Map[T, U any](arr []T, transform func(T) U) []U {
  result := make([]U, 0, len(arr))
  for _, val := range arr {
    result = append(result, transform(val))
  }
  return result
}

func main() {
  arr := []int{1, 2, 3}
  result := Map(arr, func(x int) string {
      return fmt.Sprintf("%d!", x)
  })
  fmt.Println(result) // Output: [1! 2! 3!]
}
