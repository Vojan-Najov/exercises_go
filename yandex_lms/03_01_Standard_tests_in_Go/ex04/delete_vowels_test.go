/*
 * Тест DeleteVowels
 *
 * Напишите тест для функции DeleteVowels(s string) string, которая должна удалять все
 * гласные из строки английского языка (y не считается гласной).
 * Используйте table driven testing.
 *
 * Примечания
 * Функцию DeleteVowels реализовывать не нужно.
 *
 */

package main

import (
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "test1", in: "abc", want: "bc"},
		{name: "test2", in: "abecudifovy", want: "bcdfvy"},
		{name: "test3", in: "bc", want: "bc"},
		{name: "test4", in: "ao", want: ""},
		{name: "test5", in: "couy", want: "cy"},
		{name: "test6", in: "cv d", want: "cv d"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteVowels(tc.in)
			if got != tc.want {
				t.Errorf(
					"DeleteVolwes(%q) = %q, want %q",
					tc.in,
					got,
					tc.want,
				)
			}
		})
	}
}
