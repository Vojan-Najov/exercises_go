/*
Файлы. Изменение файла

Напишите функцию ModifyFile(filename string, pos int, val string), которая
будет изменять содержимое файла на значение val, начиная с позиции pos.

Для перемещения по файлу используйте функцию os.Seek.

Примечания
Функцию main описывать не требуется.
*/

package main

import (
  "os"
  "fmt"
)

func ModifyFile(filename string, pos int, val string) {
  file, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0600)
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }
  defer file.Close()

  if _, err := file.Seek(int64(pos), 0); err != nil {
    fmt.Println("Error: ", err)
    return
  }

  file.WriteString(val)
}

func main() {
  ModifyFile("tmp.txt", 0, "123")
  ModifyFile("tmp.txt", 3, "456")
  ModifyFile("tmp.txt", 1, "00")
}
