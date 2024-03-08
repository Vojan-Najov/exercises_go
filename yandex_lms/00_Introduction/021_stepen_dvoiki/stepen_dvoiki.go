package main

import "fmt"

func IsPowerOfTwoRecursive(N int) {
  if (N == -2 || N == 2 || N == 1 || N == -1) {
    fmt.Println("YES")
    return
  }
  if N % 2 != 0 {
    fmt.Println("NO")
    return
  }
  IsPowerOfTwoRecursive(N / 2)
}

func main() {
  for {
    var n int
    if _, err := fmt.Scanln(&n); err != nil {
      fmt.Println("Некорректный ввод")
      continue
    }
    IsPowerOfTwoRecursive(n)
  }
}
