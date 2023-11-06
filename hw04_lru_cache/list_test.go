package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("once", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]

		require.Equal(t, 1, l.Len())
	})

	t.Run("four element", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushFront(20) // [20]
		l.PushFront(30) // [30]
		l.PushFront(40) // [40]

		require.Equal(t, 4, l.Len())
	})

	t.Run("remove back element", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]//4
		l.PushFront(20) // [20]//3
		l.PushFront(30) // [30]//2
		l.PushFront(40) // [40]//1

		require.Equal(t, 4, l.Len())
		require.Equal(t, 40, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)

		lastItem := l.Back()
		l.Remove(lastItem)
		require.Equal(t, 3, l.Len())
		require.Equal(t, 20, l.Back().Value)
	})

	t.Run("remove front element", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]//4
		l.PushFront(20) // [20]//3
		l.PushFront(30) // [30]//2
		l.PushFront(40) // [40]//1

		require.Equal(t, 40, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)

		firstItem := l.Front()

		l.Remove(firstItem)

		require.Equal(t, 30, l.Front().Value)
	})

	t.Run("push back element", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]//2
		l.PushFront(20) // [20]//1
		l.PushBack(30)  //[30]//3

		require.Equal(t, 3, l.Len())
		require.Equal(t, 30, l.Back().Value)

	})

	t.Run("move to front element", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]//4
		l.PushFront(20) // [20]//3
		l.PushFront(30) // [30]//2
		l.PushFront(40) // [40]//1

		require.Equal(t, 4, l.Len())
		lastElem := l.Back()
		l.MoveToFront(lastElem)
		firstElem := l.Front()
		require.Equal(t, 10, firstElem.Value)

		lastElem2 := l.Back()
		l.MoveToFront(lastElem2)
		firstElem2 := l.Front()
		require.Equal(t, 20, firstElem2.Value)

	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
