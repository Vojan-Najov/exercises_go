package main

import (
  "fmt"
  "slices"
)

func Permutations(input string) []string {
  array := []rune(input)
  slices.Sort(array)

  var strings = []string{}
  strings = append(strings, string(array))

  n := len(array)
  p := make([]int, n)
  for i := 0; i < n; i++ {
    p[i] = i
  }
  
  i := 1;
  for i < n {
    p[i]--

    j := i % 2 * p[i];
    array[j], array[i] = array[i], array[j]

    strings = append(strings, string(array))
    i = 1
    for i < n && p[i] == 0 {
      p[i] = i
      i++
    }
  }

  slices.Sort(strings)
  return strings
}

func main() {
  strings := Permutations("cba")
  fmt.Println(strings)
  strings = Permutations("a")
  fmt.Println(strings)
}
