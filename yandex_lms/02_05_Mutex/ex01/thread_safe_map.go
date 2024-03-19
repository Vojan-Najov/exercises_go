/*
1. Thread safe map

Реализуйте потокобезопасную мапу. Для чтения элементов используйте функцию
func (s *SafeMap) Get(key string) interface{}, а для записи
func (s *SafeMap) Set(key string, value interface{}) .
Используйте func NewSafeMap() *SafeMap для получению нового экземпляра.
Примечания

Код должен содержать структуру:

type SafeMap struct {
  m map[string]interface{}
  mux sync.Mutex
}
*/

package main

import (
	"sync"
)

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.m[key]
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{})}
}
