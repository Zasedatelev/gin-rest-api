package cache

import (
	"container/list"

	"sync"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRUCache struct {
	sync.RWMutex
	items    map[string]*list.Element
	queue    *list.List
	capacity int
}

func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		items:    make(map[string]*list.Element),
		queue:    list.New(),
		capacity: capacity,
	}
}

func (c *LRUCache) Set(key string, value interface{}) error {
	c.RLock()

	defer c.RUnlock()

	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element

	return nil
}

func (c *LRUCache) Get(key string) interface{} {
	c.Lock()
	defer c.Unlock()

	element, exists := c.items[key]
	if exists == false {
		return nil
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (c *LRUCache) GetLenQueue() int {
	return c.queue.Len()
}

func (c *LRUCache) GetQueue() map[string]*list.Element {
	return c.items
}
