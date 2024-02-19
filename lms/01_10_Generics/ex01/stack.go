/*
Дженерик стэк

Реализуйте дженерик структуру данных стека, которая может извлекать элементы
любого типа. Функции, которые нужно реализовать Push(val T) Pop() T

Примечания
Структура должна иметь вид Stack[T any] и иметь конструктор NewStack.
*/

package main

import "fmt"

type Stack[T any] struct {
  arr []T
  size int
}

func NewStack[T any]() *Stack[T] {
  return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
  if cap(s.arr) == s.size {
    s.arr = append(s.arr, val)
  } else {
    s.arr[s.size] = val
  }
  s.size++
}

func (s *Stack[T]) Pop() T {
  s.size--
  return s.arr[s.size]
}

func main() {
  s := NewStack[int]()
  s.Push(1)
  s.Pop()
  s.Push(1)
  s.Push(2)
  s.Push(3)
  fmt.Println(s.Pop(), s.Pop(), s.Pop())
}
