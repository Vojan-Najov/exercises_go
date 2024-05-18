/*
 * Обменник
 *
 * Напишите функцию, которая обменивает значения между двумя переменными с
 * использованием атомиков.
 * Примечания
 * Нужно реализовать для нее метод AtomicSwap, который принимает два указателя.
*/

package main

import (
	"sync/atomic"
)

func AtomicSwap(a, b *int32) {
	tmpA := atomic.LoadInt32(a)
	atomic.SwapInt32(a, atomic.LoadInt32(b))
	atomic.SwapInt32(b, tmpA)
}
