// Exercise 3.10
// Write a non-recursive version of comma, using bytes.Buffer instead of
// string concatenation.

package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if (len(s)-i-1)%3 == 0 && len(s)-i-1 != 0 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}
