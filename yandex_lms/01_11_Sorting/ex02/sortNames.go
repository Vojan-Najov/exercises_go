/*
Сортировка слайса строк

Напишите программу, содержащую функцию `SortNames(names []string)``,
которая сортирует список имён в алфавитном порядке. Если первые символы
совпадают, сортировать по вторым, и т.д.

Примечания
Пример отсортированного списка: Аксинья, Арина, Варвара, Есения

*/

package main

import (
  "fmt"
  "slices"
)

func SortNames(names []string) {
  slices.Sort(names)
}

func main() {
  names := []string{"Варвара", "Аксинья", "Есения", "Арина"}
  fmt.Println(names)
  SortNames(names)
  fmt.Println(names)
}
