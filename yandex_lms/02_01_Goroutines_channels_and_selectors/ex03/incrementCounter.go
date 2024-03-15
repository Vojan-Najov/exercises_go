/*
Мьютексы

Напишите программу, где несколько горутин func incrementCounter() одновременно
пытаются увеличить значение общей переменной с использованием мьютекса. Функция
инкремент должна вызываться несколько раз.

Примечания
var (
    counter = 0
    mu      sync.Mutex
)
*/

package main

import (
	"math/rand"
	"sync"
	"time"
)

var (
	counter = 0
	mu      sync.Mutex
)

func main() {
	go incrementCounter()
	go incrementCounter()
	time.Sleep(time.Second)
}

func incrementCounter() {
	n := rand.Intn(5)
	for i := 0; i < n; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}
