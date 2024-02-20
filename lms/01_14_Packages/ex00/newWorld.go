/*
Конструктор

Измените функцию NewWorld, чтобы она проверяла переданные размеры на
положительное число и возвращала ошибку, если это условие нарушено.

Примечания

Код программы должен содержать описание струкрутры World:
type World struct {
Height int
Width int
Cells [][]bool
}

Функцию main писать не нужно
Код поместите в пакет main

*/

package main

import (
  "fmt"
  "errors"
)

type World struct {
  Height int
  Width int
  Cells [][]bool
}

func NewWorld(height, width int) (*World, error) {
  if height < 1 {
    return nil, errors.New("non-positive height")
  }
  if width < 1 {
    return nil, errors.New("non-positive width")
  }

  cells := make([][]bool, height)
  for i := range cells {
    cells[i] = make([]bool, width)
  }

  return &World{
    Height: height,
    Width: width,
    Cells: cells,
  }, nil
}

func main() {
  w, _ := NewWorld(10, 12)
  fmt.Println(w)
  _, err := NewWorld(-10, 5)
  fmt.Println(err)
}
