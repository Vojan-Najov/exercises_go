/*
3. Thread safe queue

Напишите потокобезопасную очередь ConcurrentQueue. Реализуете следующий интерфейс:

type Queue interface {
  Enqueue(element interface{}) // положить элемент в очередь
  Dequeue() interface{} // забрать первый элемент из очереди
}

Примечания
Код должен содержать следующую структуру:

type ConcurrentQueue struct {
  queue []interface{} // здесь хранить элементы очереди
  mutex sync.Mutex
}
*/

package main

import "sync"

var _ Queue = &ConcurrentQueue{}

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(element interface{}) {
	q.mutex.Lock()
	q.queue = append(q.queue, element)
	q.mutex.Unlock()
}

func (q *ConcurrentQueue) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	element := q.queue[0]
	q.queue = q.queue[1:]
	return element
}
