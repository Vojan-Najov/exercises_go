package main
import (
  "os"
  "fmt"
  "time"
  "bufio"
  "strings"
  "errors"
  "math/rand"
)

type World struct {
  Height int
  Width int
  Cells [][]bool
}

func NewWorld(height, width int) *World {
  cells := make([][]bool, height)
  for i := range cells {
     cells[i] = make([]bool, width)
  }
  return &World{
    Height: height,
    Width: width,
    Cells: cells,
  }
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

func (w *World) NeighborsOnThor(x, y int) int {
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

func (w *World) Next(x, y int) bool {
  //n := w.Neighbors(x, y)
  n := w.NeighborsOnThor(x, y)
  alive := w.Cells[y][x]

  if n < 4 && n > 1 && alive {
    return true
  }
  if n == 3 && !alive {
    return true
  }
  return false
}

func NextState(oldWorld, newWorld *World) {
  for i := 0; i < oldWorld.Height; i++ {
    for j := 0; j < oldWorld.Width; j++ {
      newWorld.Cells[i][j] = oldWorld.Next(j, i)
    }
  }
}

func (w *World) Seed() {
  for _, row := range w.Cells {
    for i := range row {
      if rand.Intn(10) == 1 {
        row[i] = true
      }
    }
  }
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

func main() {
  height := 10
  width := 10
  currentWorld := NewWorld(height, width)
  nextWorld := NewWorld(height, width)

  currentWorld.Seed()
  for {
    fmt.Println(currentWorld)
    NextState(currentWorld, nextWorld)
    currentWorld, nextWorld = nextWorld, currentWorld
    time.Sleep(100 * time.Millisecond)
    fmt.Print("\033[H\033[2J")
  }
}

