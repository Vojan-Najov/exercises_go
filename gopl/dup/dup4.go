package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			if _, ok := counts[input.Text()]; !ok {
				counts[input.Text()] = make(map[string]bool)
			}
			counts[input.Text()][filename] = true
		}
		f.Close()
	}
	for line, m := range counts {
		if len(m) > 1 {
			var s, sep string
			for k, _ := range m {
				s += sep + k
				sep = " "
			}
			fmt.Printf("%s: %s\n", s, line)
		}
	}
}
