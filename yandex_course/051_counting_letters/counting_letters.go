package main

import (
  "strings"
  "fmt"
  "bufio"
  "os"
)

func main() {
  var s1, s2  string
  fmt.Scanln(&s1)
  fmt.Scanln(&s2)

  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')

  c1 := strings.Count(text, s1)
  c2 := strings.Count(text, s2)

  if c1 >= c2 {
    fmt.Println(s1, " ", c1)
    fmt.Println(s2, " ", c2)
  } else {
    fmt.Println(s2, " ", c2)
    fmt.Println(s1, " ", c1)
  }
}
