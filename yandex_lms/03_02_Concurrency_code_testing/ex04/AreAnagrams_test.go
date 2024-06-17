/*
 * Тест Anagrams
 *
 * Напишите тест для функции, которая проверяет, являются ли слова анограммами:
 *
 * func AreAnagrams(str1, str2 string) bool {
 *    str1 = strings.ToLower(str1)
 *    str2 = strings.ToLower(str2)
 *
 *    if len(str1) != len(str2) {
 *        return false
 *    }
 *
 *    // Convert strings to slices of runes for sorting
 *    r1 := []rune(str1)
 *    r2 := []rune(str2)
 *
 *    sort.Slice(r1, func(i, j int) bool {
 *        return r1[i] < r1[j]
 *    })
 *
 *    sort.Slice(r2, func(i, j int) bool {
 *        return r2[i] < r2[j]
 *    })
 *
 *    return string(r1) == string(r2)
 * }
 *
 * Примечания
 * Функцию AreAnagrams реализовывать не нужно.
 */

package main

import (
	"testing"
)

func TestAreAnagrams(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{name: "test1", input1: "str", input2: "rs", want: false},
		{name: "test2", input1: "str", input2: "rst", want: true},
		{name: "test3", input1: "str abc", input2: "abc str", want: true},
		{name: "test4", input1: "", input2: "", want: true},
		{name: "test5", input1: "123654", input2: "7162534", want: false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := AreAnagrams(tc.input1, tc.input2)
			if got != tc.want {
				t.Errorf(
					"AreAnagrams(%q, %q) = %v, want %v",
					tc.input1,
					tc.input2,
					got,
					tc.want,
				)
			}
		})
	}
}
