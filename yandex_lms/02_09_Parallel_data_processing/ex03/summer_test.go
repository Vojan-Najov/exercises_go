package main

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		desc           string
		nums           []int
		sum            int
		chunks         int
		expectingError bool
	}{
		{
			desc:   "short",
			nums:   []int{3, 5, 3, 6, 6},
			sum:    23,
			chunks: 2,
		},
		{
			desc:   "a bit longer",
			nums:   []int{3, 5, 3, 6, 6, 3, 5, 3, 6, 6},
			sum:    46,
			chunks: 3,
		},
		{
			desc:           "negative",
			nums:           []int{3, 5, 3, 6, 6},
			chunks:         -10,
			expectingError: true,
		},
	}
	for _, tC := range testCases {
		tt := tC
		t.Run(tC.desc, func(t *testing.T) {

			sum, err := ProcessSum(SumChunk, tt.nums, tt.chunks)
			if tt.expectingError {
				if err == nil {
					t.Errorf("expecting error, got none")
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if sum != tt.sum {
				t.Errorf("expecting sum: %v, got sum: %v:", tt.sum, sum)
			}
		})
	}
}

type MockSummer struct {
	called int
}

func (ms *MockSummer) sum(arr []int, result chan<- int) {
	ms.called++
}

func TestChunkcs(t *testing.T) {
	testCases := []struct {
		desc      string
		nums      []int
		chunkSize int
	}{
		{
			desc:      "short",
			nums:      []int{3, 5, 3, 6, 6},
			chunkSize: 2,
		},
		{
			desc:      "longer",
			nums:      []int{3, 5, 3, 6, 6, 3, 5, 3, 6, 6},
			chunkSize: 4,
		},
		{
			desc:      "empty",
			nums:      []int{},
			chunkSize: 10,
		},
	}

	for _, tC := range testCases {
		tt := tC
		t.Run(tC.desc, func(t *testing.T) {

			ms := MockSummer{}
			_, err := ProcessSum(ms.sum, tt.nums, tt.chunkSize)

			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			numberOfChunks := math.Ceil(float64(len(tt.nums)) / float64(tt.chunkSize))

			if ms.called != int(numberOfChunks) {
				t.Errorf("expecting to call summer: %v times, but was called: %v:", numberOfChunks, ms.called)
			}
		})
	}
}
