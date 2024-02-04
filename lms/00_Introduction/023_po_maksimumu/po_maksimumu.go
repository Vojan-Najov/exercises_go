package main

import "fmt"

func FindMinMaxInSlice(slice []int) (int, int) {
  if len(slice) == 0 {
    return 0, 0
  }
  var mini int = slice[0]
  var maxi int = mini
  for i := 1; i < len(slice); i++ {
    if slice[i] > maxi {
      maxi = slice[i]
    }
    if slice[i] < mini {
      mini = slice[i]
    }
  }
  return mini, maxi
}

func main() {
  slice := make([]int, 100)
  for i := 50; i >= 0; i-- {
    slice[i] = i
  }
  for i := 51; i < 100; i++ {
    slice[i] = i
  }

  mini, maxi := FindMinMaxInSlice(slice)
  fmt.Println(mini, maxi)
}
