/*
Аргументы

Напишите функцию run() error, которая будет считывать параметры из
командой строки: размер сетки и процент изначального заполнения.
Пример запуска приложения main с заданными параметрами:
go build -o main // скомпилируем программу
main 10 20 30 // сетка 10x20, процент заполнения - 30
Полученные значения запишите в файл config.txt в формате: 10x20 30%
*/

package main

import (
  "os"
  "fmt"
  "errors"
  "strconv"
)

func run() error {
  filename := "config.txt"

  args := os.Args[1:]
  if len(args) < 3 {
    return errors.New("Incorrect number of arguments")
  }

  width, err := strconv.Atoi(args[0])
  if err != nil || width <= 0{
    return errors.New("Incorrect width")
  }
  height, err := strconv.Atoi(args[1])
  if err != nil || height <= 0 {
    return errors.New("Incorrect height")
  }
  percent, err := strconv.Atoi(args[2])
  if err != nil || percent <= 0 {
    return errors.New("Incorrect percentage")
  }

  f, err := os.OpenFile(filename,
                        os.O_CREATE | os.O_TRUNC | os.O_WRONLY,
                        0644)
  if err != nil {
    return err
  }
  defer f.Close()

  fmt.Fprintf(f, "%dx%d %d%%\n",
              width, height, percent)

  return nil
}

func main() {
  err := run()
  if err != nil {
    fmt.Println(err)
  }
}
