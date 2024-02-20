/*
Сортировка символов по частоте

Дана строка с символами из набора алфавита. Напишите программу с функцией
SortByFreq(str string) string, которая будет сортировать символы из строки
по возрастанию с учетом частоты повторения. Символы с наименьшим количеством
вхождений должны идти в начале, а символы с наибольшей частотой - в конце.\
В случае одинаковой частоты символов, они должны быть отсортированы в
алфавитном порядке.

Примечания

Пример:
Вход: "abbbzzzat"
Выход: "taabbbzzz"
*/

package main

import (
  "fmt"
  "sort"
  "slices"
)

func SortByFreq(str string) string {
  runes := []rune(str)
  slices.Sort(runes)

  sl := make([][]rune, 0, len(runes))
  i := 0
  for i < len(runes) {
    j := i + 1
    for j < len(runes) && runes[j] == runes[i] {
      j++
    }
    sl = append(sl, runes[i:j])
    i = j
  }
  sort.Slice(sl, func (i, j int) bool {
    return len(sl[i]) < len(sl[j]) || 
          len(sl[i]) == len(sl[j]) && sl[i][0] < sl[j][0]
  })

  var result string  
  for _, r := range sl {
    result += string(r)
  }

  return result
}

func main() {
  str := "abbbzzzat"
  fmt.Println(str)
  str = SortByFreq(str)
  fmt.Println(str)
}
