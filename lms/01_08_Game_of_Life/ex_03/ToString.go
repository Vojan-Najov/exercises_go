/*
Вывод на экран

В приведенном коде игры текущее состояние сетки выводится на экран
следующим образом: fmt.Println(currentWorld), что не совсем наглядно.
Чтобы получить наглядное изображение нашей сетки мы можем отрисовывать
клетки разного цвета, например:
brownSquare := "\xF0\x9F\x9F\xAB"
greenSquare := "\xF0\x9F\x9F\xA9"

Для этого, созданный нами тип type World struct {
Height int // высота сетки
Width int // ширина сетки
Cells [][]bool
}
должен иметь метод String(), который будет формировать строку для
отображения. Напишите данный метод и попробуйте запустить всю программу,
поэкспериментируйте с разными символами и цветами.

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

func (w *World) String() string {
  count := 4
  symbols := make([]byte, w.Height * w.Width * count + w.Height - 1)

  var i int
  for k, row := range w.Cells {
    for _, cell := range row {
      if cell {
        symbols[i] = '\xF0'
        i++
        symbols[i] = '\x9F'
        i++
        symbols[i] = '\x9F'
        i++
        symbols[i] = '\xA9'
        i++
      } else {
        symbols[i] = '\xF0'
        i++
        symbols[i] = '\x9F'
        i++
        symbols[i] = '\x9F'
        i++
        symbols[i] = '\xAB'
        i++
      }
    }
    if k < w.Height - 1 {
      symbols[i] = '\n'
      i++
    }
  }

  return string(symbols)
}
