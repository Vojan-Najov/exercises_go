/*
 * Тест Sort
 *
 * Функция SortIntegers(numbers []int) (пакет main) сортирует переданный слайс в порядке
 * увеличения значений. Напишите тест для проверки корректности работы.
 * Примечания
 * Функцию SortIntegers реализовывать не нужно.
 */

package main

import (
	"slices"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "test1", input: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "test2", input: []int{4, 5, 1, 0, 3}, want: []int{0, 1, 3, 4, 5}},
		{name: "test3", input: []int{4}, want: []int{4}},
		{name: "test4", input: []int{}, want: []int{}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := make([]int, len(tc.input))
			copy(got, tc.input)
			SortIntegers(got)

			if !slices.Equal(got, tc.want) {
				t.Errorf(
					"SortIntegers(%v), got: %v, want: %v",
					tc.input,
					got,
					tc.want,
				)
			}
		})
	}
}
