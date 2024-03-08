/*
Подсчёт соседей

Напишите недостающий метод для подсчета живых соседних клеток в сетке:
func (w *World) Neighbors(x, y int) int.
Всего у клетки может быть максимум восемь соседей. Например, клетка в
центре (зеленые клетки - живые) имеет три соседа.
example image

Примечания
Код программы должен содержать описание струкрутры World:
type World struct { Height int Width int Cells [][]bool }
*/

package main

type World struct {
  Height int
  Width int
  Cells [][]bool
}

func (w *World) Neighbors(x, y int) int {
  var n int
  y_prev := (y - 1 + w.Height) % w.Height
  x_prev := (x - 1 + w.Width) % w.Width
  y_next := (y + 1) % w.Height
  x_next := (x + 1) % w.Width

  if w.Cells[y_prev][x_prev] {
    n++
  }
  if w.Cells[y_prev][x] {
    n++
  }
  if w.Cells[y_prev][x_next] {
    n++
  }
  if w.Cells[y][x_prev] {
    n++
  }
  if w.Cells[y][x_next] {
    n++
  }
  if w.Cells[y_next][x_prev] {
    n++
  }
  if w.Cells[y_next][x] {
    n++
  }
  if w.Cells[y_next][x_next] {
    n++
  }

  return n
}
