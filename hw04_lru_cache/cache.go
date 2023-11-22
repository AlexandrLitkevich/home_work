package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	mu       sync.Mutex
	items    map[Key]*ListItem
}

type element struct {
	Key   Key
	Value any
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	elem := element{Value: value, Key: key}
	item, ok := l.items[key]

	if !ok {
		l.items[key] = l.queue.PushFront(elem)
	} else {
		item.Value = elem
		l.queue.MoveToFront(item)
	}

	if l.queue.Len() > l.capacity {
		lastItem := l.queue.Back()
		delete(l.items, lastItem.Value.(element).Key)
		l.queue.Remove(lastItem)
	}

	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	item, ok := l.items[key]
	if !ok {
		return nil, false
	}
	l.queue.MoveToFront(item)
	return l.items[key].Value.(element).Value, true
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.capacity = 0
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
