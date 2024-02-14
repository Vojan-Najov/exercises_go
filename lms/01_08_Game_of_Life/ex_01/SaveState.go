/*
Сохранение состояния

Напишите новый метод
func (w *World) SaveState(filename string) error
для сохранения текущего состояния сетки в файл.
Метод должен создавать новый файл и записать данные в бинарном виде,
например:
110011
100101
Примечания

Код программы должен содержать описание струкрутры World:
type World struct { Height int Width int Cells [][]bool }
*/

package main

import (
  "os"
  "fmt"
  "math/rand"
)

type World struct {
  Height int
  Width int
  Cells [][]bool
}

func (w *World) SaveState(filename string) error {
  file, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer file.Close()

  for i := 0; i < w.Height; i++ {
    for j := 0; j < w.Width; j++ {
      if w.Cells[i][j] {
        fmt.Fprint(file, 1)
      } else {
        fmt.Fprint(file, 0)
      }
    }
    if i < w.Height - 1 {
      fmt.Fprintln(file)
    }
  }

  return nil
}

func main() {
  height := 5
  width := 10
  cells := make([][]bool, height)
  for i := range cells {
    cells[i] = make([]bool, width)
  }
  for _, row := range cells {
    for i := range row {
      if rand.Intn(100) % 2 == 0 {
        row[i] = true
      }
    }
  }

  world := World{Height: 5, Width: 10, Cells: cells}
  if err := world.SaveState("state.txt"); err != nil {
    fmt.Println(err)
  }
}
