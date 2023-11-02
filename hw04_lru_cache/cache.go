package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
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

		l.queue.Remove(lastItem)
		delete(l.items, lastItem.Key)
	}

	return true
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := l.items[key]; ok {
		return item.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
