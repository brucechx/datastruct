package double_list

import (
	"testing"
	"fmt"
)

func TestList(t *testing.T)  {
	l := New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	printListInfo(l)
	fmt.Println(l.Front().Value)

}

func printListInfo(l *List){
	for e := l.Front(); e != nil; e = e.Next(){
		fmt.Print(e.Value," ")
	}
	fmt.Println()
}
