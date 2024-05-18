package main

import (
	"context"
	"slices"
	"sync"
	"testing"
)

type Num struct {
	int
}

// getSequence реализует интерфейс sequenced.
func (n Num) getSequence() int {
	// для упрощения примера будем считать
	// само число как порядковый номер части
	return n.int
}

func (n Num) Ordered() {

}

func TestProcessTempCh(t *testing.T) {
	testCases := []struct {
		desc  string
		input [][]Num
		out   []Num
	}{
		{
			desc: "",
			input: [][]Num{
				{{0}, {2}},
				{{1}, {3}},
			},
			out: []Num{{0}, {1}, {2}, {3}},
		},
		{
			desc: "",
			input: [][]Num{
				{{0}, {2}},
				{{1}},
			},
			out: []Num{{0}, {1}, {2}},
		},
		{
			desc: "",
			input: [][]Num{
				{{0}, {1}},
				{},
			},
			out: []Num{{0}, {1}},
		},
		{
			desc: "",
			input: [][]Num{
				{{0}, {4}},
				{{1}, {5}},
				{{2}, {6}},
				{{3}},
			},
			out: []Num{{0}, {1}, {2}, {3}, {4}, {5}, {6}},
		},
	}
	for _, tC := range testCases {
		tt := tC
		t.Run(tt.desc, func(t *testing.T) {

			ctx := context.Background()

			tempChan := make(chan fanInRecord[Num])

			go func() {
				var wg sync.WaitGroup

				for index, sl := range tt.input {
					wg.Add(1)
					index := index
					sl := sl
					go func() {
						defer wg.Done()
						pauseChan := make(chan struct{})
						for _, v := range sl {
							record := fanInRecord[Num]{
								index: index,
								data:  v,
								pause: pauseChan,
							}
							tempChan <- record

							<-record.pause

						}
					}()

				}
				wg.Wait()
				close(tempChan)
			}()

			outChan := processTempCh(ctx, 10, tempChan)

			out := []Num{}

			for v := range outChan {
				out = append(out, v)
			}

			if !slices.Equal(out, tt.out) {
				t.Errorf("expected out: %v, got: %v", tt.out, out)
			}

		})
	}
}
