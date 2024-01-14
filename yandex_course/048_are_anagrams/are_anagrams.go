package main

import (
  "fmt"
  "strings"
  "slices"
)

func AreAnagrams(str1, str2 string) bool {
  str1 = strings.ReplaceAll(str1, " ", "") 
  str2 = strings.ReplaceAll(str2, " ", "") 
  s1 := []rune(strings.ToLower(str1))
  s2 := []rune(strings.ToLower(str2))
  slices.Sort(s1)
  slices.Sort(s2)
  return string(s1) == string(s2)
}

func main() {
  fmt.Println(AreAnagrams("ab", "ba"))
  fmt.Println(AreAnagrams("ab", "b  a"))
  fmt.Println(AreAnagrams("aB", "b  A"))
  fmt.Println(AreAnagrams("cab", "b  a"))
  fmt.Println(AreAnagrams("Кабан", "банка"))
}
