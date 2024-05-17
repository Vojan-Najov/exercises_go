/*
 * 2. Cache patterns
 *
 * MRU (most recently used) cache. В случае если при установке элемента
 * достигнуто максимальное значение размера кеша, то удаляем из кеша последний
 * использованный элемент.
 *
 * Название структуры и сигнатура функций определены ниже.
 *
 * // структура MRU кеша
 * type MRUCache struct {}
 *
 * // возвращает новый инстанс кеша размером capacity
 * func NewMRUCache(capacity int) *MRUCache
 * // устанавливает значени value ключу key
 * func (c *MRUCache) Set(key, value string)
 * // получает значение и флаг его начличия по ключу key
 * func (c *MRUCache) Get(key string) (string, bool)
 *
 */

package main

import (
	"container/list"
	"fmt"
	"sync"
)

type MRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mutex    sync.Mutex
}

type CacheEntry struct {
	key   string
	value string
}

func NewMRUCache(capacity int) *MRUCache {
	return &MRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (c *MRUCache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, found := c.cache[key]; found {
		element.Value.(*CacheEntry).value = value
		c.list.MoveToFront(element)
	} else {
		entry := &CacheEntry{key: key, value: value}
		element := c.list.PushFront(entry)
		c.cache[key] = element

		if c.list.Len() > c.capacity {
			newest := c.list.Back()
			if newest != nil {
				delete(c.cache, newest.Value.(*CacheEntry).key)
				c.list.Remove(newest)
			}
		}

	}
}

func (c *MRUCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, found := c.cache[key]; found {
		c.list.MoveToFront(element)
		return element.Value.(*CacheEntry).value, true
	}

	return "", false
}

func (c *MRUCache) PrintCache() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fmt.Printf(
		"MRU Cache (Capacity: %d, Size: %d): [",
		c.capacity,
		c.list.Len(),
	)
	for element := c.list.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*CacheEntry)
		fmt.Printf("(%v: %v) ", entry.key, entry.value)
	}
	fmt.Println("]")
}
