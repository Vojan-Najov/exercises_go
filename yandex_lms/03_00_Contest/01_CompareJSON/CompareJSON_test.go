/*
 * Тест CompareJSON
 *
 * Функция CompareJSON(json1, json2 []byte) (bool, error), принимает на вход два объекта
 * json и сравнивает их. Если они равны, то функция возвращает true, nil,
 * иначе false, nil, либо описание ошибки.
 * Напишите тесты для проверки корректности работы.
 * Примечания
 *
 *    порядок следования полей в json в равных объектах может быть разный
 *    json не содержит вложенных объектов
 *    покрытие кода функции должно быть 100%.
 */

package main

import (
	"errors"
	"testing"
)

const (
	jsonStr1 = `{"name": "John", "age": 30, "Gender": "male"}`
	jsonStr2 = `{"name": "John", "age": 33, "Gender": "male"}`
	jsonStr3 = `{"age": 30, "name": "John", "Gender": "male"}`
	jsonStr4 = `{"name": "John", "age": 30, some_error "Gender": "male"}`
)

func TestCompareJSON(t *testing.T) {
	cases := []struct {
		name string
		in1  string
		in2  string
		want bool
		err  error
	}{
		{name: "test1", in1: "", in2: "", want: false, err: errors.New("tmp")},
		{name: "test2", in1: jsonStr1, in2: jsonStr1, want: true, err: nil},
		{name: "test3", in1: jsonStr1, in2: jsonStr2, want: false, err: nil},
		{name: "test4", in1: jsonStr1, in2: jsonStr3, want: true, err: nil},
		{name: "test4", in1: jsonStr2, in2: jsonStr3, want: false, err: nil},
		{
			name: "test5",
			in1:  jsonStr1,
			in2:  jsonStr4,
			want: false,
			err:  errors.New("tmp"),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := CompareJSON([]byte(tc.in1), []byte(tc.in2))
			if got != tc.want || (err != nil && tc.err == nil) ||
				(tc.err != nil && err == nil) {
				t.Errorf(
					"CompareJSON(%v, %v) = %v, %v, want %v, %v",
					string(tc.in1),
					string(tc.in2),
					got,
					err,
					tc.want,
					tc.err,
				)
			}
		})
	}
}
