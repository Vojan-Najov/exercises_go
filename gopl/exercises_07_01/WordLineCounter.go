// Exercise 7.1:
// Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.

package main

import (
	"bufio"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var words int
	for {
		advance, _, err := bufio.ScanWords(p, true)
		if err != nil {
			return 0, nil
		}
		if advance == 0 {
			break
		}
		words++
		p = p[advance:]
	}
	*c += WordCounter(words)
	return words, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	var lines int
	for {
		advance, _, err := bufio.ScanLines(p, true)
		if err != nil {
			return 0, nil
		}
		if advance == 0 {
			break
		}
		lines++
		p = p[advance:]
	}
	*c += LineCounter(lines)
	return lines, nil
}

func main() {
	var wc WordCounter
	fmt.Fprintf(&wc, "Hello, my friend %s %s", "Joe", "Doe")
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, "hi\n My name is Joji\n\nMe good\n")
	fmt.Println(lc)
}
