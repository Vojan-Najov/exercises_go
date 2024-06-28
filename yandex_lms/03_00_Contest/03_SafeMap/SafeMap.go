package main

import (
	"sync"
)

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()

	return s.m[key]
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.m[key] = value
}

/*
// TestConcurrentCache - тест для параллельного тестирования кеша
func TestConcurrentCache(t *testing.T) {
	// Создаем новый кеш
	cache := NewCache()

	// Добавляем запись в кеш (не в параллели)
	cache.Add("key1", "value1")

	// Запускаем тесты в параллели
	t.Run("AddRecord", func(t *testing.T) {
		t.Parallel()
		// Добавляем запись в кеш
		cache.Add("key2", "value2")
	})

	t.Run("DeleteRecord", func(t *testing.T) {
		t.Parallel()
		// Удаляем запись из кеша
		cache.Delete("key1")
	})

	// Проверяем наличие записи после всех операций
	t.Run("CheckRecord", func(t *testing.T) {
		t.Parallel()
		// Получаем значение записи по ключу
		value := cache.Get("key2")
		expectedValue := "value2"

		// Проверяем, что значение соответствует ожидаемому
		if value != expectedValue {
			t.Errorf("Expected value: %s, Actual value: %s", expectedValue, value)
		}
	})
}
*/
