/*
 * Lock-free очередь *
 *
 * Реализуйте структуру - lock-free очередь, где операции добавления и
 * извлечения элементов производятся без использования блокировок, а с
 * использованием CAS (Compare and Swap) на указателях.
 *
 * Примечания
 * Сами структуры выглядят таким образом.
 * type Node struct {
 *   value int
 *   next  *Node
 * }
 *
 * type LockFreeQueue struct {
 *    head unsafe.Pointer
 *    tail unsafe.Pointer
 * }
 *
 * Нужно реализовать для них методы.
 * NewLockFreeQueue() *LockFreeQueue
 * Enqueue(value int)
 * Dequeue() (int, bool)
 *
 */

package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type Node struct {
	value int
	next  unsafe.Pointer
}

type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	dummy := &Node{value: 0, next: nil}
	return &LockFreeQueue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}

}

func (q *LockFreeQueue) Dequeue() (int, bool) {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		nextHead := (*Node)(head).next

		if head == tail {
			if nextHead == nil {
				return 0, false
			} else {
				atomic.CompareAndSwapPointer(
					&q.tail,
					tail,
					nextHead,
				)
			}
		} else {
			result := (*Node)(nextHead).value
			if atomic.CompareAndSwapPointer(
				&q.head,
				head,
				nextHead,
			) {
				return result, true
			}
		}
	}
}

func (q *LockFreeQueue) Enqueue(value int) {
	newTail := &Node{value: value}
	for {
		tail := atomic.LoadPointer(&q.tail)
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(&(*Node)(tail).next),
			nil,
			unsafe.Pointer(newTail),
		) {
			atomic.CompareAndSwapPointer(
				&q.tail,
				tail,
				unsafe.Pointer(newTail),
			)
			return
		} else {
			atomic.CompareAndSwapPointer(
				&q.tail,
				tail,
				unsafe.Pointer((*Node)(tail).next),
			)
		}
	}
}

func (q *LockFreeQueue) Print() {
	// Загружаем голову списка
	curr := atomic.LoadPointer(&q.head)

	// Проходим по всем узлам в списке
	fmt.Print("{ ")
	for curr != nil {
		// Преобразуем указатель в структуру Node
		node := (*Node)(curr)
		// Выводим значение узла
		fmt.Print(node.value, " ")
		// Загружаем указатель на следующий узел
		curr = atomic.LoadPointer(&node.next)
	}
	fmt.Println("}")

}
