package impl

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node
	val  interface{}
}

type List struct {
	head Node
	size int
}

func NewList() *List {
	return &List{}
}

func (*List) At(i int) interface{} {
	panic("implement me")
}

func (l *List) Insert(idx int, data interface{}) error {
	if l == nil {
		return errors.New("list is null")
	}
	if idx > l.size {
		return errors.New("cap too small")
	} else {
		cur := &l.head
		for i := 0; i < idx; i++ {
			cur = cur.next
		}
		t := cur.next
		cur.next = &Node{val: data}
		t.next = cur.next.next
		l.size++
	}
	return nil
}

func (*List) InsertFront(data interface{}) error {
	panic("implement me")
}

func (l *List) InsertEnd(data interface{}) error {
	if l == nil {
		l = NewList()
	}
	cur := &l.head
	for i := 0; i < l.size; i++ {
		cur = cur.next
	}
	cur.next = &Node{}
	cur.val = data
	l.size++
	return nil
}

func (*List) Remove(i int) error {
	panic("implement me")
}

func (*List) RemoveFront() error {
	panic("implement me")
}

func (*List) RemoveEnd() error {
	panic("implement me")
}

//func (l *List) Next() bool {
//	return l.cur.next == nil
//}
//
//func (l *List) reset() {
//	l.cur = l.head.next
//}
//
//func (l *List) Data() interface{} {
//	return l.cur.val
//}

func (l *List) Output() {
	cur := &l.head
	for i := 0; i < l.size; i++ {
		fmt.Println("output: ", cur.val)
		cur = cur.next
	}

}
