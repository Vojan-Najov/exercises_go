/*
 * Напишите тест для функции:
 *
 * func ReverseString(input string) string {
 *     runes := []rune(input)
 *     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
 *          runes[i], runes[j] = runes[j], runes[i]
 *     }
 *     return string(runes)
 * }
 *
 * Примечания
 * Функцию ReverseString реализовывать не нужно.
 */

package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "test1", input: "abc", want: "cba"},
		{name: "test2", input: "aaa", want: "aaa"},
		{name: "test3", input: "aaaa", want: "aaaa"},
		{name: "test4", input: "1234567", want: "7654321"},
		{name: "test5", input: "12345678", want: "87654321"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := ReverseString(tc.input)
			if got != tc.want {
				t.Errorf(
					"Reverse(%v) = %v, want %v",
					tc.input,
					got,
					tc.want,
				)
			}
		})
	}
}
