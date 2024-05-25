package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	fmt.Printf("%d\t%d\t%v\n", len(x), cap(x), x)

	z := x[3:7]
	fmt.Printf("%d\t%d\t%v\n", len(z), cap(z), z)

	z = appendInt(z, 0)
	fmt.Printf("%d\t%d\t%v\n", len(z), cap(z), z)
	fmt.Printf("%d\t%d\t%v\n", len(x), cap(x), x)
}
