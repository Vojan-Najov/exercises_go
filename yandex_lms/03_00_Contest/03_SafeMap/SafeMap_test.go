/*
 * Тест SafeMap
 *
 * Протестируйте следующую реализацию потокобезопасной мапы:
 *
 * type SafeMap struct {
 *     m   map[string]interface{}
 *    mux sync.Mutex
 * }
 *
 * func NewSafeMap() *SafeMap {
 *     return &SafeMap{
 *         m: make(map[string]interface{}),
 *     }
 * }
 *
 * func (s *SafeMap) Get(key string) interface{} {
 *     s.mux.Lock()
 *     defer s.mux.Unlock()
 *
 *     return s.m[key]
 * }
 *
 * func (s *SafeMap) Set(key string, value interface{}) {
 *    s.mux.Lock()
 *    defer s.mux.Unlock()
 *
 *    s.m[key] = value
 * }
 *
 * Примечания
 *   покрытие кода тестами должно быть 100%.
 *
 */

package main

import (
	"testing"
)

func TestSafeMap(t *testing.T) {
	m := NewSafeMap()

	m.Set("1", 1)
	m.Set("2", 2)

	t.Run("SetKey", func(t *testing.T) {
		t.Parallel()
		m.Set("3", 3)
	})

	t.Run("GetKey", func(t *testing.T) {
		t.Parallel()
		value, ok := m.Get("1").(int)
		expectedValue := 1

		if !ok || value != expectedValue {
			t.Errorf(
				"Expected value: %d, Actual value: %v %T",
				expectedValue,
				value,
				value,
			)
		}
	})
}
