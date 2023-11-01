package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
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
	return &ListItem{}
}

func (l *list) Back() *ListItem {
	return &ListItem{}
}

func (l *list) PushFront(v interface{}) *ListItem {
	return &ListItem{}
}

func (l *list) PushBack(v interface{}) *ListItem {
	return &ListItem{}
}

func (l *list) Remove(i *ListItem) {}

func (l *list) MoveToFront(i *ListItem) {}

func NewList() List {
	return &list{}
}
