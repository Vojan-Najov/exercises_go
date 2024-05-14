package main

import (
	"context"
	"slices"
	"testing"
)

func TestFanInInt(t *testing.T) {
	testCases := []struct {
		desc        string
		input       [][]int
		expectedOut []int
	}{
		{
			desc: "",
			input: [][]int{
				{3, 5, 6},
				{1, 7, 9},
				{10, 11, 12},
			},
			expectedOut: []int{1, 3, 5, 6, 7, 9, 10, 11, 12},
		},
		{
			desc: "",
			input: [][]int{
				{},
				{1, 2},
				{5, 10},
			},
			expectedOut: []int{1, 2, 10, 5},
		},
	}
	for _, tC := range testCases {
		tt := tC

		t.Run(tt.desc, func(t *testing.T) {

			inputChannels := []<-chan int{}

			for _, v := range tt.input {
				v := v
				inputChan := make(chan int)
				inputChannels = append(inputChannels, inputChan)
				go func() {

					for _, val := range v {
						inputChan <- val
					}
					close(inputChan)
				}()
			}

			ctx := context.Background()
			outChan := FanIn[int](ctx, inputChannels...)

			joined := []int{}

			for v := range outChan {
				joined = append(joined, v)
			}
			slices.Sort(tt.expectedOut)
			slices.Sort(joined)

			if !slices.Equal(tt.expectedOut, joined) {
				t.Errorf("expected out: %v, got : %v", tt.expectedOut, joined)
			}

		})
	}
}

func TestFanInString(t *testing.T) {
	testCases := []struct {
		desc        string
		input       [][]string
		expectedOut []string
	}{
		{
			desc: "",
			input: [][]string{
				{"3", "5", "6"},
				{"1", "7", "9"},
				{"10", "11", "12"},
			},
			expectedOut: []string{"1", "3", "5", "6", "7", "9", "10", "11", "12"},
		},
		{
			desc: "",
			input: [][]string{
				{},
				{"", "2"},
				{"5", "10"},
			},
			expectedOut: []string{"", "2", "10", "5"},
		},
	}
	for _, tC := range testCases {
		tt := tC

		t.Run(tt.desc, func(t *testing.T) {

			inputChannels := []<-chan string{}

			for _, v := range tt.input {
				v := v
				inputChan := make(chan string)
				inputChannels = append(inputChannels, inputChan)
				go func() {

					for _, val := range v {
						inputChan <- val
					}
					close(inputChan)
				}()
			}

			ctx := context.Background()
			outChan := FanIn[string](ctx, inputChannels...)

			joined := []string{}

			for v := range outChan {
				joined = append(joined, v)
			}

			slices.Sort(tt.expectedOut)
			slices.Sort(joined)

			if !slices.Equal(tt.expectedOut, joined) {
				t.Errorf("expected out: %v, got : %v", tt.expectedOut, joined)
			}

		})
	}
}
