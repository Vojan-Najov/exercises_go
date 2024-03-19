/*
4. Thread safe buf

Напишите функцию Write(num int), которая записывает данные в буфер Buf []int.
Функция Consume() int должна забирать первое значение из этого буфера и возвращать
его. Используйте мьютекс для синхронизации доступа к буферу.
*/

package main

import "sync"

var (
	Buf []int
	mu  sync.Mutex
)

func Write(num int) {
	mu.Lock()
	Buf = append(Buf, num)
	mu.Unlock()
}

func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	num := Buf[0]
	Buf = Buf[1:]
	return num
}
