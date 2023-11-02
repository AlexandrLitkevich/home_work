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

// https://cs.opensource.google/go/go/+/refs/tags/go1.21.3:src/container/list/list.go;l=230
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

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	fmt.Println("l.tail", l.tail)
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item, _ := v.(ListItem)
	fmt.Println("this item", item.Value)
	fmt.Println("this item", v)

	l.len++
	//if l.head == nil {
	//	l.head = item
	//	l.tail = item
	//	fmt.Println("this l.head", l.head)
	//	fmt.Println("this l.tail", l.tail)
	//	return item
	//}
	//
	//l.head = item
	//item.Prev = l.head.Prev
	//item.Next = l.tail
	return &item
}

func (l *list) PushBack(value interface{}) *ListItem {
	return &ListItem{}
}

func (l *list) Remove(item *ListItem) {
	//item.Next.Prev = item.Next //Я удаляю всегда последний
	item.Prev.Next = item.Prev
}

func (l *list) MoveToFront(i *ListItem) {}

func NewList() List {
	return &list{}
}
