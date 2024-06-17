/*
 * Тест Sum
 *
 * Функция Sum(a, b int) int (пакет main) возвращает результат суммирования чисел a и b.
 * Напишите тест TestSum(t *testing.T) для проверки корректности работы.
 *
 * Примечания
 * Функцию Sum(a, b int) int реализовывать не нужно.
 */

package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "test1",
			values: []int{1, 1},
			want:   2,
		},
		{
			name:   "test2",
			values: []int{0, 0},
			want:   0,
		},
		{
			name:   "test3",
			values: []int{-1, 1},
			want:   0,
		},
		{
			name:   "test4",
			values: []int{19, 11},
			want:   30,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.values[0], tc.values[1])
			if got != tc.want {
				t.Errorf(
					"Sum(%v, %v) = %v; want %v",
					tc.values[0],
					tc.values[1],
					got,
					tc.want,
				)
			}
		})
	}
}
