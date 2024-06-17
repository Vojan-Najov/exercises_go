/*
 * Тест GetUTFLength
 *
 * Фукнция GetUTFLength(input []byte) (int, error) возвращает длину строки UTF8 и ошибку
 * ErrInvalidUTF8 (в случае возникновения). Напишите тест, который бы проверял
 * возвращаемые функцией значения.
 *
 * var ErrInvalidUTF8 = errors.New("invalid utf8")
 *
 * func GetUTFLength(input []byte) (int, error) {
 *   if !utf8.Valid(input) {
 *     return 0, ErrInvalidUTF8
 *  }
 *
 *  return utf8.RuneCount(input), nil
 * }
 *
 * Примечания
 * Функцию GetUTFLength реализовывать не нужно.
 */

package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
		err   error
	}{
		{name: "test1", input: "abc", want: 3, err: nil},
		{name: "test2", input: "abc абц", want: 7, err: nil},
		{name: "test3", input: "", want: 0, err: nil},
		{name: "test3", input: "\xd0", want: 0, err: ErrInvalidUTF8},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetUTFLength([]byte(tc.input))
			if got != tc.want || err != tc.err {
				t.Errorf(
					"GetUTFLength(%v) = %v, %v, but want %v, %v",
					[]byte(tc.input),
					got,
					err,
					tc.want,
					tc.err,
				)
			}
		})
	}
}
