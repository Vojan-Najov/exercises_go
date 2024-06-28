/*
 * NumbersGen
 *
 * Напишите тесты для следующего кода:
 *
 * func NumbersGen(filename string) <-chan int {
 *     output := make(chan int)
 *
 *     go func() {
 *        defer close(output)
 *
 *        file, err := os.Open(filename)
 *        if err != nil {
 *             return
 *        }
 *        defer file.Close()
 *
 *        scanner := bufio.NewScanner(file)
 *        for scanner.Scan() {
 *            line := scanner.Text()
 *            num, err := strconv.Atoi(line)
 *            if err == nil {
 *                output <- num
 *            }
 *        }
 *    }()
 *
 *    return output
 * }
 *
 * Примечания
 *   покрытие кода функции должно быть 100%.
 */

package main

import (
	"os"
	"testing"
)

func TestNumbersGenNoExist(t *testing.T) {
	output := NumbersGen("noexist")
	_, ok := <-output
	if ok {
		t.Errorf("Expected empty closed chain")
	}
}

const numbers = `
123
34
45
ad
7
`

func TestNumbersGen(t *testing.T) {
	filename := "tmp.txt"
	err := os.WriteFile(filename, []byte(numbers), 0644)
	if err != nil {
		t.Errorf("err")
		return
	}
	expectedValues := []int{123, 34, 45, 7}

	output := NumbersGen(filename)

	for _, expectedValue := range expectedValues {
		value, ok := <-output
		if !ok || value != expectedValue {
			t.Errorf("Expected value: %d, but got: %d", expectedValue, value)
		}
	}
	_, ok := <-output
	if ok {
		t.Errorf("Expected closed empty chain")
	}

	os.RemoveAll(filename)
}
