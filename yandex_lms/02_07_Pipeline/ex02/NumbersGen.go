/*
 * 2. Генератор чисел
 *
 * Напишите функцию NumbersGen(filename string) <-chan int, которая будет
 * читать файл и записывать в выходной канал числа из этого файла.
 * Формат файла: каждая строка содержит одно число, например:
 * 2
 * 4
 * 6
 *
 * Примечания
 * Если в строке не число, пропускайте значение.
 */

package main

import (
	"bufio"
	"os"
	"strconv"
)

func NumbersGen(filename string) <-chan int {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}

	out := make(chan int)
	go func() {
		defer f.Close()
		defer close(out)
		fs := bufio.NewScanner(f)
		for fs.Scan() {
			num, err := strconv.Atoi(fs.Text())
			if err != nil {
				continue
			}
			out <- num
		}
	}()

	return out
}
