package main

import (
	"bufio"
	"os"
	"strconv"
)

func NumbersGen(filename string) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			num, err := strconv.Atoi(line)
			if err == nil {
				output <- num
			}
		}
	}()

	return output
}
