package hw04lrucache

import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	GetMap() map[Key]*ListItem
}

type lruCache struct {
	//mu sync.Mutex
	capacity int //размер очереди
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	// TODO  added mutex
	if _, ok := l.items[key]; !ok {
		item := &ListItem{Value: value, Key: key}
		l.items[key] = item
		l.queue.PushFront(item)
		return true
	}

	l.items[key].Value = value

	if l.capacity == l.queue.Len() {
		lastItem := l.queue.Back()
		//TODO Remove without MAP
		l.queue.Remove(lastItem)
		fmt.Println("this lastItem.key", lastItem.Key)
		delete(l.items, lastItem.Key)
	}

	return true
}

// Get DONE
func (l *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

// TDOD
func (l *lruCache) Clear() {}

func (l *lruCache) GetMap() map[Key]*ListItem {
	return l.items
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
