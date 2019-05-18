package impl

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewList(t *testing.T) {
	l := new(List)
	t.Log("nil l ", l == nil, l)
}

func TestList_InsertEnd(t *testing.T) {
	l := new(List)
	for i := 10; i < 15; i++ {
		l.InsertEnd(i)
	}
	fmt.Println(strings.Repeat("===", 20))
	l.Output()
}

func TestList_Insert(t *testing.T) {
	l := new(List)
	//for i := 0; i < 15; i++ {
	l.Insert(0, 10)
	//}
	fmt.Println(strings.Repeat("===", 20))
	l.Output()
}
