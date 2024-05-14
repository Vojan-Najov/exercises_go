/*
 * Параллельная обработка Summer
 *
 * Напишите программу для подсчета суммы чисел в слайсе параллельно, используя
 * горутины. Программа должна разделять исходный слайс на несколько слайсов и
 * затем находить сумму каждой части в отдельной горутине. После этого
 * подсчитывается финальная сумма. Программа должна реализовать следующий
 * интерфейс:
 * type Summer interface{
 *   // функция для нахождения суммы чисел
 *   func ProcessSum(
 *     // функция, которая будет вызываться для нахождения суммы части слайса.
 *     // результат суммы записать в result
 *     summer func(arr []int, result chan<- int),
 *     // слайс с числами, сумму которых нужно найти
 *     nums []int,
 *     // сколько элементов в одной части (последняя часть может быть меньше)
 *     сhunkSize int,
 *   ) (int, error) // вернуть сумму чисел
 * }
 * Также нужно реализовать функцию SumChunk(arr []int, result chan<- int),
 * которая будет вызываться для нахождения суммы части слайса (summer).
 * В случае возникновения ошибок верните 0 и возникшую ошибку.
 *
 */

package main

import (
	"errors"
	"sync"
)

func ProcessSum(
	summer func(arr []int, result chan<- int),
	nums []int,
	chunkSize int,
) (int, error) {
	if chunkSize <= 0 {
		return 0, errors.New("non-positive chunk's size")
	}

	size := len(nums) / chunkSize
	if len(nums)%chunkSize != 0 {
		size++
	}

	wg := sync.WaitGroup{}
	wg.Add(size)
	result := make(chan int, size)
	defer close(result)

	var first, last int = 0, chunkSize
	if last > len(nums) {
		last = len(nums)
	}

	for first <= len(nums) {
		go func(arr []int) {
			defer wg.Done()
			summer(arr, result)
		}(nums[first:last])

		first += chunkSize
		last += chunkSize
		if last > len(nums) {
			last = len(nums)
		}
	}

	wg.Wait()

	var sum int
	for i := 0; i < size; i++ {
		select {
		case v, ok := <-result:
			if ok {
				sum += v
			} else {
				break
			}
		default:
			break
		}
	}

	return sum, nil
}

func SumChunk(arr []int, result chan<- int) {
	var sum int
	for _, v := range arr {
		sum += v
	}
	result <- sum
}
