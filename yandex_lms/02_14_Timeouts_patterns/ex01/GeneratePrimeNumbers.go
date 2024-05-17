/*
 * 1. timeouts_patterns
 *
 * Напишите программу для генерации простых чисел вплоть до N (2 < N < 105) и
 * используйте для реализации паттерна таймаута. Выводите каждое простое число
 * на новой строке. Через 0,01 секунду и остановите генерацию простых чисел.
 * Реализуйте функцию
 * func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int)
 * stop - канал для остановки генерации
 * prime_nums - канал для вывода простых чисел
 * N - число до которого нужно генерировать числа
 */

package main

import (
	"time"
)

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	primes := make(map[int]struct{})

	time.AfterFunc(10*time.Millisecond, func() {
		stop <- struct{}{}
	})
	for n := 2; n < N; n++ {
		isPrime := true
		for k, _ := range primes {
			if n%k == 0 {
				isPrime = false
				break
			}
		}
		select {
		case <-stop:
			return
		default:
			if isPrime {
				primes[n] = struct{}{}
				prime_nums <- n
			}
		}
	}
	close(prime_nums)
}
