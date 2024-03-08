/*
Пишем в файл

Напишите функцию WriteToLogFile(message string, fileName string) error,
которая пишет в файл output.txt строку "hello world"

*/

package main

import (
  "os"
  "log"
)

func WriteToLogFile(message string, fileName string) error {
  file, err := os.OpenFile(fileName,
                           os.O_CREATE | os.O_WRONLY | os.O_APPEND,
                           0644)
  if err != nil {
    return err
  }
  defer file.Close()

  log.SetOutput(file)
  log.Println(message)

  return nil
}

func main() {
  WriteToLogFile("hello world", "tmp.log")
}
