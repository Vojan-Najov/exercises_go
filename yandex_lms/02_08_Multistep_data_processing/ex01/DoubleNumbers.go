/*
 * 1. Удвоение чисел
 *
 * Напишите функцию
 * DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int,
 * чтобы она удваивала элементы из канала in и записывала их в выходной канал.
 * Функция должна завершать работу при закрытии канала done, либо при закрытии
 * канала in.
 */

package main

func DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
				return
			default:
				out <- (2 * v)
			}
		}
	}()
	return out
}
