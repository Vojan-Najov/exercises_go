package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err)
		return
	}

	fib1 := 0
	fib2 := 1
	for fib2 < n {
		fib2 += fib1
		fib1 = fib2 - fib1
	}
	for i := 0; i < 10; i++ {
		fmt.Println(fib2)
		fib2 += fib1
		fib1 = fib2 - fib1
	}
}
