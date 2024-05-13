/*
 * 3. Пайплайн - сумма чётных чисел
 *
 * Реализуйте pipeline, который читает файл с числами (каждая строка - число),
 * игнорирует нечетные числа, а чётные суммирует.
 * Для этого напишите функцию SumValuesPipeline(filename string) int,
 * вызывающую:
 * - NumbersGen(filename string) <-chan int для чтения файла
 * - Filter(in <-chan int) <-chan int - для фильтрации чисел
 * - Sum(in <-chan int) int - для суммирования результата.
 * Напишите реализацию перечисленных функций
 */

package main

import (
	"bufio"
	"os"
	"strconv"
)

func SumValuesPipeline(filename string) int {
	numbers := NumbersGen(filename)
	filter_numbers := Filter(numbers)
	return Sum(filter_numbers)
}

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

func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num%2 == 0 {
				out <- num
			}
		}
	}()
	return out
}

func Sum(in <-chan int) int {
	var sum int
	for num := range in {
		sum += num
	}
	return sum
}
