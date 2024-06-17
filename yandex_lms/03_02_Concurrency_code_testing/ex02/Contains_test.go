/*
 * Тест Contains numbers
 *
 * Напишите тест для функции:
 *
 * func Contains(numbers []int, target int) bool{
 *    for _, num := range numbers {
 *        if num == target {
 *            return true
 *        }
 *    }
 *    return false
 * }
 *
 * Примечания
 *
 * Функцию Contains реализовывать не нужно.
 */
package main

import (
	"sync"
	"testing"
)

func TestContains(t *testing.T) {
	var wg sync.WaitGroup

	cases := []struct {
		name   string
		input  []int
		target int
		want   bool
	}{
		{name: "test1", input: []int{1, 2, 3}, target: 1, want: true},
		{name: "test2", input: []int{1, 2, 3}, target: 4, want: false},
		{name: "test3", input: []int{1, 2, 3, 2}, target: 2, want: true},
		{name: "test4", input: []int{}, target: 2, want: false},
	}

	wg.Add(len(cases))
	for i := 0; i < len(cases); i++ {
		go func(i int) {
			defer wg.Done()
			got := Contains(cases[i].input, cases[i].target)
			if got != cases[i].want {
				t.Errorf(
					"Contains(%v, %v) = %v, want: %v",
					cases[i].input,
					cases[i].target,
					got,
					cases[i].want,
				)
			}
		}(i)
	}

	wg.Wait()
}
