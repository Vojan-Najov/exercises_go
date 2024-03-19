package main

import (
	"sync"
	"testing"
)

func TestProducerConsumer(t *testing.T) {

	Buf = []int{}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			Write(i)
		}()
	}

	wg.Wait()

	if len(Buf) != 1000 {
		t.Errorf("Unexpected len. Got: %d, Expected: %d", len(Buf), 1000)
	}

	var results []int

	var resultMu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			res := Consume()

			resultMu.Lock()
			results = append(results, res)
			resultMu.Unlock()
		}()
	}
	wg.Wait()

	if len(Buf) != 0 {
		t.Errorf("Unexpected len. Got: %d, Expected: %d", len(Buf), 0)
	}
	if len(results) != 1000 {
		t.Errorf("Unexpected results len. Got: %d, Expected: %d", len(results), 1000)
	}

}
