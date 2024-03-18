/*
Исправьте еще одну гонку

Исправьте гонку в коде ниже

func SafeCounterFunc(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        *counter++
    }
}
*/

package main

import "sync"

func SafeCounterFunc(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}
