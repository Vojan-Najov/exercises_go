// echo5 prints the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%3d: %s\n", i, arg)
	}
}
