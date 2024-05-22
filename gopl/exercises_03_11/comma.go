// Exercise 3.11
// Enhance comma so that it deals correctly with floating-point numbers and an
// optional sign.

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	end := strings.Index(s, ".")
	if end == -1 {
		end = len(s)
	}

	var buf bytes.Buffer
	buf.WriteByte('[')
	var i int = 0
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		i++
	}
	for i < end {
		buf.WriteByte(s[i])
		if (end-i-1)%3 == 0 && end-i-1 != 0 {
			buf.WriteByte(',')
		}
		i++
	}
	for i < len(s) {
		buf.WriteByte(s[i])
		i++
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

	fmt.Println(comma("-1"))
	fmt.Println(comma("-12"))
	fmt.Println(comma("-123"))
	fmt.Println(comma("-1234"))
	fmt.Println(comma("-12345"))
	fmt.Println(comma("-123456"))
	fmt.Println(comma("-1234567"))

	fmt.Println(comma("1.0987"))
	fmt.Println(comma("12.0987"))
	fmt.Println(comma("123.0987"))
	fmt.Println(comma("1234.0987"))
	fmt.Println(comma("12345.0987"))
	fmt.Println(comma("123456.0987"))
	fmt.Println(comma("1234567.0987"))

	fmt.Println(comma("-1.0987"))
	fmt.Println(comma("-12.0987"))
	fmt.Println(comma("-123.0987"))
	fmt.Println(comma("-1234.0987"))
	fmt.Println(comma("-12345.0987"))
	fmt.Println(comma("-123456.0987"))
	fmt.Println(comma("-1234567.0987"))
}
