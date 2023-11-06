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
	item := &ListItem{
		Value: v,
	}

	l.len++
	if l.head == nil {
		l.head = item
		l.tail = item
		return item
	}

	item.Prev = l.head.Prev
	l.head.Prev = item
	item.Next = l.head
	l.head = item

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
	}

	switch {
	case l.tail == nil:
		l.head = item
		l.tail = item
	default:
		item.Prev = l.tail
		item.Prev.Next = item
		l.tail = item
	}
	l.len++

	return item
}

func (l *list) Remove(item *ListItem) {
	defer func() {
		l.len--
	}()

	switch {
	case item.Next == nil:
		item.Prev.Next = nil
		l.tail = item.Prev
		return
	case item.Prev == nil:
		item.Next.Prev = nil
		l.head = item.Next
		return
	default:
		item.Prev.Next = item.Next
		item.Next.Prev = item.Prev
		return
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	if i.Next == nil {
		l.tail = i.Prev
	}
	i.Prev.Next = i.Next
	i.Prev = nil
	i.Next = l.head
	l.head = i
}

func NewList() List {
	return &list{}
}
