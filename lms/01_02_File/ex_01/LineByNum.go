/*
Файлы. Чтение строки по номеру

Напишите функцию LineByNum(inputFilename string, lineNum int) string,
которая получает в качестве параметров имя файла и номер строки, а
возвращает текст строки по ее порядковому номеру в файле (нумерация с нуля).

Если строки с указанным номером найти не удается, то верните пустую строку.

Примечания
Функцию main создавать не надо.
Нумерация элементов в программировании всегда начинается с 0.
*/

package main

import (
  "os"
  "io"
  "fmt"
  "bufio"
)

func LineByNum(inputFilename string, lineNum int) string {
  file, err := os.Open(inputFilename)
  defer file.Close()

  if err != nil || lineNum < 0 {
    return ""
  }

  fileScanner := bufio.NewScanner(file)
  for fileScanner.Scan() {
    if lineNum == 0 {
      break
    }
    lineNum--
  }
  return fileScanner.Text()
}

func main() {
  for {
    var num int
    _, err := fmt.Scanln(&num)
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Println("Error: ", err)
      break
    }

    fmt.Println(LineByNum("LineByNum.go", num))
  }
}
