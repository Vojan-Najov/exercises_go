package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err)
		return
	}

	for i := 3; i <= n; i += 3 {
		fmt.Println(i)
	}
}
