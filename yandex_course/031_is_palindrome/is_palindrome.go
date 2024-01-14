package main

import (
  "fmt"
  "strings"
)

func IsPalindrome(input string) bool {
  array := []rune(strings.ToLower(strings.Replace(input, " ", "", -1)))
  size := len(array)
  for idx, r := range array {
    if r != array[size - 1 - idx] {
      return false
    }
  }
  return true
}

func main() {
  fmt.Println(IsPalindrome("А роза упала на лапу Азора"))
  fmt.Println(IsPalindrome("abceba"))
  fmt.Println(IsPalindrome("abcba"))
}
