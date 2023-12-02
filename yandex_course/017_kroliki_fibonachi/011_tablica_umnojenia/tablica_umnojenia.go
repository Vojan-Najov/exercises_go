package main

import "fmt"

func main() {
    var n int
    if _, err := fmt.Scanln(&n); err != nil {
        fmt.Println(err)
		return
    }
    for i := 1; i <= 10; i++ {
        fmt.Println(n, "*", i, "=", n * i)
    }
}
