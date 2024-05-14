package main

import (
	"context"
	"slices"
	"testing"
)

type Num struct {
	int
}

func (n Num) getSequence() int {
	return n.int
}

func TestEvenNumbersGen(t *testing.T) {
	testCases := []struct {
		desc    string
		nums    []Num
		outNums []Num
	}{
		{
			desc: "",
			nums: []Num{
				{1}, {2}, {3}, {4}, {5}, {6},
			},
			outNums: []Num{
				{2}, {4}, {6},
			},
		},
		{
			desc: "",
			nums: []Num{
				{-1}, {-2}, {3}, {4}, {5}, {6},
			},
			outNums: []Num{
				{-2}, {4}, {6},
			},
		},
		{
			desc: "",
			nums: []Num{
				{10}, {11}, {12}, {13}, {14}, {15}, {16},
			},
			outNums: []Num{
				{10}, {12}, {14}, {16},
			},
		},
		{
			desc:    "",
			nums:    []Num{},
			outNums: []Num{},
		},
		{
			desc:    "",
			nums:    []Num{{11}},
			outNums: []Num{},
		},
		{
			desc:    "",
			nums:    []Num{{12}},
			outNums: []Num{{12}},
		},
	}
	for _, tC := range testCases {
		tt := tC
		t.Run(tt.desc, func(t *testing.T) {
			ctx := context.Background()

			inputCh1 := EvenNumbersGen(ctx, tt.nums...)

			out := []Num{}

			for v := range inputCh1 {
				out = append(out, v)
			}

			if !slices.Equal(out, tt.outNums) {
				t.Errorf("expected out: %v, got: %v", tt.outNums, out)
			}
		})
	}
}
