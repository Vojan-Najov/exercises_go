package main

import (
  "fmt"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
  runes := []rune(str)
  if position >= len(runes) {
    return rune(0), fmt.Errorf("position out of range")
  }
  return runes[position], nil
}

func main() {
  if r, err := GetCharacterAtPosition("Абв", 1); err == nil {
    fmt.Println(string(r))
  } else {
    fmt.Println(err)
  }
  if r, err := GetCharacterAtPosition("Абв", 4); err == nil {
    fmt.Println(string(r))
  } else {
    fmt.Println(err)
  }
}
