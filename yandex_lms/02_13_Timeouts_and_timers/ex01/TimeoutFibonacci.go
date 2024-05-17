/*
 * 1. timeouts
 *
 * Сгенерируйте N-е число Фибоначчи. Установите ограничение в T и, если оно
 * пройдет в течение указанного времени, выведите число Фибоначчи в одной строке
 * без кавычек “Fibonacci(N) = D”, где D - N-е число Фибоначчи, в противном
 * случае выдайте ошибку.
 * Сигнатура функции
 * TimeoutFibonacci(n int, timeout time.Duration) (int, error)
 * n - номер числа
 * timeout - время, отведенное на операцию
 *
 */

package main

import (
	"errors"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	c := make(chan int)
	go func() {
		fib1, fib2 := 0, 1
		for n > 0 {
			fib1, fib2 = fib2, fib1+fib2
			n--
		}
		c <- fib1
	}()

	select {
	case res := <-c:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}
