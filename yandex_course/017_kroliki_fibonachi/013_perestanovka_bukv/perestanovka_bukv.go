package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err)
		return
	}

	prod := 1
	for i := n; i > 1; i-- {
		prod *= i
	}
	fmt.Println(prod)
}
