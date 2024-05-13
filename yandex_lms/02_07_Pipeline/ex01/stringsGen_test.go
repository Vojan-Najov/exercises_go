package main

import (
	"testing"
)

func TestStringsGen(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			[]string{"Hello", "World"},
			[]string{"Hello", "World"},
		},
		{
			[]string{"Golang", "is", "awesome"},
			[]string{"Golang", "is", "awesome"},
		},
		{
			[]string{},
			[]string{},
		},
	}

	for _, tc := range testCases {
		gen := StringsGen(tc.input...)
		i := 0
		for line := range gen {
			if line != tc.expected[i] {
				t.Errorf(
					"Expected %s, but got %s",
					tc.expected[i],
					line,
				)
			}
			i++
		}
	}
}
