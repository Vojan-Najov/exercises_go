package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
