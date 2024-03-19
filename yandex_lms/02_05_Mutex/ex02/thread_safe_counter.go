/*
2. Thread safe counter

Напишите потокобезопасный счётчик Counter. Реализуете следующий интерфейс:

type Сount interface{
  Increment() // увеличение счётчика на единицу
  GetValue() int // получение текущего значения
}

Примечания
Код должен содержать следующую структуру:

type Counter struct {
  value int // значение счетчика
  mu sync.RWMutex
}
*/

package main

import "sync"

type Counter struct {
	value int
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
