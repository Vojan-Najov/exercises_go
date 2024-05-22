// Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sumFlag = flag.Int("sha", 256, "256, 384 or 512 for sha")

func init() {
	flag.Parse()
	fmt.Println(*sumFlag)
	if *sumFlag != 256 && *sumFlag != 384 && *sumFlag != 512 {
		fmt.Fprintln(os.Stderr, "incorrect sha")
		os.Exit(1)
	}
}

func main() {

	for _, arg := range flag.Args() {
		switch *sumFlag {
		case 256:
			fmt.Printf("%x\n", sha256.Sum256([]byte(arg)))
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(arg)))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(arg)))
		}
	}
}
