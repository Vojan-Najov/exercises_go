package main

import "fmt"

func SwapKeysAndValues(m map[string]string) map[string]string {
  var swapped = map[string]string{}
  for key, value := range m {
    swapped[value] = key
  }
  return swapped
}

func main() {
  var m = map[string]string{"1": "a", "2": "b"}
  fmt.Println(m)
  swapped := SwapKeysAndValues(m)
  fmt.Println(swapped)
}
