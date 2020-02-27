package array_stack

import (
	"bytes"
	"fmt"
)

// 用slice 实现
// 先进后出
type Item interface {}

type Stack interface {
	Push(i Item)
	Pop() Item
	IsEmpty() bool
	Size() int
}

type ArrayStack struct {
	Items []Item
}

func NewStack() *ArrayStack{
	return &ArrayStack{
		Items: []Item{},
	}
}

// push
func (s *ArrayStack)Push(i Item)  {
	s.Items = append(s.Items, i)
}

// pop
func (s *ArrayStack)Pop() Item {
	if s.IsEmpty(){
		panic("栈为空")
	}
	item := s.Items[len(s.Items) - 1]
	s.Items = s.Items[:len(s.Items) - 1]
	return item
}

// isEmpty
func (s *ArrayStack)IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *ArrayStack)Size() int {
	return len(s.Items)
}

// 遍历显示队列中的所有元素
func (s *ArrayStack)String() string{
	var buffer bytes.Buffer
	for i:=s.Size()-1; i>=0; i--{
		buffer.WriteString(fmt.Sprintf("%v", s.Items[i]))
	}
	return buffer.String()
}
