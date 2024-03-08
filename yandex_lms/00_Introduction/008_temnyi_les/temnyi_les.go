package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	switch {
	case num < 0 && -num%2 == 0:
		fmt.Println("Число отрицательное и четное")
	case num < 0 && -num%2 == 1:
		fmt.Println("Число отрицательное и нечетное")
	case num > 0 && num%2 == 0:
		fmt.Println("Число положительное и четное")
	case num > 0 && num%2 == 1:
		fmt.Println("Число положительное и нечетное")
	}
}
