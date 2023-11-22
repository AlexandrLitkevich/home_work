package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(item *ListItem)
	MoveToFront(item *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{v, l.head, nil}
	if l.head == nil {
		l.tail = node
	} else {
		l.head.Prev = node
	}
	l.len++
	l.head = node
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{v, nil, l.tail}
	if l.tail == nil {
		l.head = node
	} else {
		l.tail.Next = node
	}
	l.len++
	l.tail = node
	return l.tail
}

func (l *list) Remove(item *ListItem) {
	if item.Prev == nil {
		l.head = item.Next
	} else {
		item.Prev.Next = item.Next
	}

	if item.Next == nil {
		l.tail = item.Prev
	} else {
		item.Next.Prev = item.Prev
	}
	l.len--
}

func (l *list) MoveToFront(item *ListItem) {
	l.PushFront(item.Value)
	l.Remove(item)
}

func NewList() List {
	return &list{}
}
