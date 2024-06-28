package main

import (
	"cmp"
	_ "encoding/json"
	"errors"
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
