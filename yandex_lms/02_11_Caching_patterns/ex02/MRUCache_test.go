package main

import (
	"testing"
)

func TestMRUCache(t *testing.T) {
	cache := NewMRUCache(2)
	cache.PrintCache()

	// Тест добавления элементов в кэш
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.PrintCache()

	// Тест обновления значения элемента в кэше
	cache.Set("key1", "new value1")
	cache.PrintCache()
	value, ok := cache.Get("key1")
	cache.PrintCache()
	if !ok || value != "new value1" {
		t.Errorf("Expected value for key1 to be 'new value1', but got '%s'", value)
	}

	// Тест получения элемента из кэша
	value, ok = cache.Get("key2")
	cache.PrintCache()
	if !ok || value != "value2" {
		t.Errorf("Expected value for key2 to be 'value2', but got '%s'", value)
	}

	// Тест получения элемента из кэша
	value, ok = cache.Get("key2")
	cache.PrintCache()
	if !ok || value != "value2" {
		t.Errorf("Expected value for key2 to be 'value2', but got '%s'", value)
	}

	// Тест получения элемента из кэша
	value, ok = cache.Get("key2")
	cache.PrintCache()
	if !ok || value != "value2" {
		t.Errorf("Expected value for key2 to be 'value2', but got '%s'", value)
	}

	// Тест получения элемента из кэша
	value, ok = cache.Get("key2")
	cache.PrintCache()
	if !ok || value != "value2" {
		t.Errorf("Expected value for key2 to be 'value2', but got '%s'", value)
	}

	// Тест заполнения кэша до максимальной ёмкости
	cache.Set("key3", "value3")
	cache.PrintCache()
	_, ok = cache.Get("key1")
	cache.PrintCache()
	if ok {
		t.Errorf("Expected key1 to be deleted from the cache")
	}

	value, ok = cache.Get("key3")
	if !ok || value != "value3" {
		t.Errorf("Expected value for key3 to be 'value3', but got '%s'", value)
	}

	value, ok = cache.Get("key2")
	if !ok || value != "value2" {
		t.Errorf("Expected value for key2 to be 'value2', but got '%s'", value)
	}
}
