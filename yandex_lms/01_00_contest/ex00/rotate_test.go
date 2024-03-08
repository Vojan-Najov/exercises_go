package main

import (
	"slices"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		data     []int
		pos      int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{10, 20, 30, 40, 50}, -2, []int{30, 40, 50, 10, 20}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 24, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{10, 20, 30, 40, 50}, 0, []int{10, 20, 30, 40, 50}},
	}

	for _, tc := range testCases {
		rotated := Rotate(tc.data, tc.pos)
		if !slices.Equal(rotated, tc.expected) {
			t.Errorf("Rotate(%v, %d) = %v, expected %v", tc.data, tc.pos, rotated, tc.expected)
		}
	}
}
