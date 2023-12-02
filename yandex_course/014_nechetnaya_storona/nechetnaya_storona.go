package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil || n < 0 {
		fmt.Println("Некорректный ввод")
		return
	}

	sum := 0
	for i := 1; i <= n; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}
