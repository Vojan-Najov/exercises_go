/*
JSON

Напишите функцию CompareJSON(json1, json2 []byte) (bool, error), которая принимает
на входе два объекта json и сравнивает их. Если они равны, то функция должна вернуть
true, nil, иначе false, nil, либо описание ошибки.

Примечания
  порядок следования полей в json в равных объектах может быть разный
  json не содержит вложенных объектов
*/

package main

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func CompareJSON(json1, json2 []byte) (bool, error) {
	fields1 := strings.Split(strings.Trim(string(json1), "{}"), ",")
	fields2 := strings.Split(strings.Trim(string(json2), "{}"), ",")

	type KeyValue struct {
		key   string
		value string
	}

	var keysValues1 []KeyValue
	for _, field := range fields1 {
		fields := strings.Split(field, ":")
		if len(fields) != 2 {
			return false, errors.New("error compare")
		}
		keysValues1 = append(
			keysValues1,
			KeyValue{
				strings.TrimSpace(fields[0]),
				strings.TrimSpace(fields[1]),
			},
		)
	}
	var keysValues2 []KeyValue
	for _, field := range fields2 {
		fields := strings.Split(field, ":")
		if len(fields) != 2 {
			return false, errors.New("error compare")
		}
		keysValues2 = append(
			keysValues2,
			KeyValue{
				strings.TrimSpace(fields[0]),
				strings.TrimSpace(fields[1]),
			},
		)
	}

	slices.SortFunc(keysValues1, func(a, b KeyValue) int {
		return cmp.Compare(a.key, b.key)
	})
	slices.SortFunc(keysValues2, func(a, b KeyValue) int {
		return cmp.Compare(a.key, b.key)
	})

	return reflect.DeepEqual(keysValues1, keysValues2), nil
}

func main() {
	type ColorGroup struct {
		ID   int
		Name string
		//Colors []string
	}
	group := ColorGroup{
		ID:   1,
		Name: "Reds",
		//Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b1, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	type ColorGroupRev struct {
		Name string
		ID   int
	}
	grouprev := ColorGroupRev{
		ID:   1,
		Name: "Reds",
	}
	b2, err := json.Marshal(grouprev)
	if err != nil {
		fmt.Println("error:", err)
	}

	ok, _ := CompareJSON(b1, b2)
	fmt.Println(ok)

	group = ColorGroup{
		ID:   2,
		Name: "Reds",
	}
	b1, err = json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	ok, _ = CompareJSON(b1, b2)
	fmt.Println(ok)

	ok, err = CompareJSON(
		[]byte(`{"name": "John", "age": 30}`),
		[]byte(`{"age": 30,"name": "John"}`),
	)
	fmt.Println(ok, err)
	ok, err = CompareJSON(
		[]byte(`{"name": "Alice", "age": 25}`),
		[]byte(`{"name": "Bob", "age": 28}`),
	)
	fmt.Println(ok, err)
	ok, err = CompareJSON(
		[]byte(`{"name": "John", "age": "30"}`),
		[]byte(`{"name": "John", "age": 30}`),
	)
	fmt.Println(ok, err)

	ok, err = CompareJSON(
		[]byte(`{"name": "John", "age": "30"}`),
		[]byte(`{"name": "John", "age" 30}`),
	)
	fmt.Println(ok, err)
}
