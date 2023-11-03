package hw04lrucache

import "fmt"

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
	Key   Key
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

// Len DONE
func (l *list) Len() int {
	return l.len
}

// Front DONE
func (l *list) Front() *ListItem {
	return l.head
}

// Back DONE
func (l *list) Back() *ListItem {
	return l.tail
}

// PushFront DONE
func (l *list) PushFront(v interface{}) *ListItem {
	el, ok := v.(*ListItem)
	fmt.Println(ok)
	if !ok {
		fmt.Println(";dklfjdkasljf")
	}
	fmt.Println("this el.Key ", el.Key)

	item := &ListItem{
		Value: el.Value,
		Key:   el.Key,
	}

	l.len++
	if l.head == nil {
		l.head = item
		l.tail = item
		return item
	}

	item.Prev = l.head.Prev //nil
	l.head.Prev = item      //?
	item.Next = l.head
	l.head = item

	return item
}

// PushBack DONE
func (l *list) PushBack(v interface{}) *ListItem {
	el, ok := v.(*ListItem)
	fmt.Println(ok)
	if !ok {
		fmt.Println(";dklfjdkasljf")
	}
	fmt.Println("this el.Key ", el.Key)

	item := &ListItem{
		Value: el.Value,
		Key:   el.Key,
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

// Remove DONE
func (l *list) Remove(item *ListItem) {
	switch {
	case item.Next == nil:
		item.Prev.Next = nil
		l.tail = item.Prev
	case item.Prev == nil:
		item.Next.Prev = nil
		l.head = item.Next
	default:
		item.Prev.Next = item.Next
		item.Next.Prev = item.Prev
	}
	l.len--

	return
}

// MoveToFront DONE
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
