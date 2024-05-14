/*
 * Синхронизация Генератор чисел
 *
 * Дан интерфейс:
 *
 * type sequenced interface {
 *   getSequence() int
 * }
 *
 * Напишите функцию
 * func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T,
 * которая получает на вход элементы типа T, а возвращает канал, который выдает
 * только те элементы из списка, для которых getSequence возвращает четное
 * число.
 * Код решения должен содержать объявление интерфейса sequenced.
 *
 */

package main

import (
	"context"
)

type sequenced interface {
	getSequence() int
}

func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for _, num := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				if num.getSequence()%2 == 0 {
					out <- num
				}
			}
		}
	}()

	return out
}
