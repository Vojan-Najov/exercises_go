/*
Исправить гонку

Исправьте race condition в коде ниже:

func RaceConditionFunc(counter *int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        *counter++
    }
}
*/

package main

import "sync"

var mu sync.Mutex

func RaceConditionFunc(counter *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}
