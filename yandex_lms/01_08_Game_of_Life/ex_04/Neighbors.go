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
  if y - 1 >= 0 {
    if x - 1 >= 0 && w.Cells[y-1][x-1] {
      n++
    }
    if w.Cells[y-1][x] {
      n++
    }
    if x + 1 < w.Width && w.Cells[y-1][x+1] {
      n++
    }
  }
  if x - 1 >= 0 && w.Cells[y][x-1] {
    n++
  }
  if x + 1 < w.Width && w.Cells[y][x+1] {
    n++
  }
  if y + 1 < w.Height {
    if x - 1 >= 0 && w.Cells[y+1][x-1] {
      n++
    }
    if w.Cells[y+1][x] {
      n++
    }
    if x + 1 < w.Width && w.Cells[y+1][x+1] {
      n++
    }
  }

  return n
}
