/*
Загрузка состояния

Напишите новый метод
func (w *World) LoadState(filename string) error
для структуры World, который будет считывать из файла исходное состояние и
устанавливать размерность сетки в соответствии с данным файлом.
Файл должен содержать бинарные данные, например:
110011
100101
где 1 - живая клетка, 0 - мертвая.
Размерность после чтения данного файла: width: 6, height: 2.
В случае нарушения размерности в файле (разное количество элементов в
строках) должна возвращаться ошибка.

Примечания
Код программы должен содержать описание струкрутры World:
type World struct { Height int Width int Cells [][]bool }
*/

package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "errors"
)

type World struct {
  Height int
  Width int
  Cells [][]bool
}

func (w *World) LoadState(filename string) error {
  var width int
  var height int
  var lines []string
  var cells [][]bool

  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer file.Close()

  fileScanner := bufio.NewScanner(file)
  for fileScanner.Scan() {
    str := strings.TrimSpace(fileScanner.Text())
    if width != 0 && len(str) != width {
      return errors.New("Incorrect file") 
    } else {
      width = len(str)
    }
    lines = append(lines, str)
  }

  height = len(lines)
  if height < 1 || width < 1 {
    return errors.New("Incorrect file") 
  }

  cells = make([][]bool, height)
  for i := 0; i < height; i++ {
    cells[i] = make([]bool, width)
  }

  for i, line := range lines {
    for j, c := range line {
      if c == '1' {
        cells[i][j] = true
      } else if c != '0' {
        return errors.New("Incorrect file")
      }
    }
  }

  w.Height = height
  w.Width = width
  w.Cells = cells

  return nil
}

func main() {
  world := World{}
  world.LoadState("state.txt")
  for i := 0; i < world.Height; i++ {
    for j := 0; j < world.Width; j++ {
      if world.Cells[i][j] {
        fmt.Print(1)
      } else {
        fmt.Print(0)
      }
    }
    fmt.Println()
  }

  err := world.LoadState("error.txt")
  if err != nil {
    fmt.Println(err)
  }
}
