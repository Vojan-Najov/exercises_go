/*
Мьютексы и мапы

Создайте программу, в которой несколько горутин параллельно обращаются к общей мапе,
выполняя операции
записи func writeToMap(key, value int)
и чтения func readFromMap(key int).
Используйте мьютексы для безопасного доступа к мапе.

Примечания

var (
    mu    sync.Mutex
    myMap = make(map[int]int)
)
*/

package main

import (
	"math/rand"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	myMap = make(map[int]int)
)

func main() {
	n := rand.Intn(10)
	for i := 0; i < n; i++ {
		go writeToMap(i, 10*n*i)
		go readFromMap(i)
		go readFromMap(i + 1)
	}
	time.Sleep(2 * time.Second)
}

func writeToMap(key, value int) {
	mu.Lock()
	myMap[key] = value
	mu.Unlock()
}

func readFromMap(key int) {
	mu.Lock()
	_ = myMap[key]
	mu.Unlock()
}
