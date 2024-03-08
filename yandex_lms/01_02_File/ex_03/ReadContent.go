/*
Файлы. Чтение

Напишите функцию ReadContent(filename string) string, которая принимает на
вход путь к файлу, а возвращает его содержимое.
В случае любой ошибки возвращайте пустую строку.

Примечания
Функцию main описывать не требуется.
*/

package main

import (
  "os"
  "fmt"
)

func ReadContent(filename string) string {
  data, err := os.ReadFile(filename)
  if err != nil {
    return ""
  }
  return string(data)
}

func main() {
  fmt.Print(ReadContent("ReadContent.go"))
}
