package link_stack

import (
	"bytes"
	"fmt"
)

// 链表实现
// 先入后出

type Item interface {}

type Stack interface {
	Push(i Item)
	Pop() Item
	IsEmpty() bool
	Size() int
}

type Node struct {
	Value Item
	pre *Node
	next *Node
}

type LinkStack struct {
	head *Node
	tail *Node
	size int
}

//新建队列，即创建Node指针head，用来指向链表第一个结点，初始为空
func NewStack() *LinkStack{
	return &LinkStack{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Push(e Element)    //向栈顶添加元素,加在头前
func (s *LinkStack) Push(i Item){
	newNode := &Node{
		Value: i,
		pre:nil,
		next:s.head,
	}
	if s.IsEmpty(){
		s.head = newNode
		s.tail = newNode
	}else {
		s.head.pre = newNode
		s.head = newNode
	}
	s.size ++
}

// isEmpty
func (s *LinkStack) IsEmpty() bool{
	return s.size == 0
}

//	Pop()   Element    //移除栈顶元素
func (s *LinkStack)Pop()  interface{}{
	if s.IsEmpty(){
		panic("栈为空")
	}
	firstNode := s.head
	node := firstNode.Value
	s.head = firstNode.next
	firstNode.next = nil
	firstNode.pre = nil
	firstNode.Value = nil
	s.size --
	return node
}

//	Top()   Element   //获取栈顶元素（不删除）
func (s *LinkStack) Top() interface{}{
	if s.IsEmpty(){
		panic("栈为空")
	}
	return s.head.Value
}

//	Clear()  bool       //清空栈
func (s *LinkStack)Clear()  {
	if s.IsEmpty(){
		panic("stack is empty")
	}
	p := s.head
	for p != nil{
		n := p.next
		p.next = nil
		p.pre = nil
		p.Value = nil
		p = n
	}
	s.head = nil
	s.tail = nil
	s.size = 0
}

//	Size()  int            //获取栈的元素个数
func (s *LinkStack)Size() int  {
	return s.size
}

// 遍历显示队列中所有元素
func (s *LinkStack)String()  string{
	if s.IsEmpty(){
		return ""
	}
	var buffer bytes.Buffer
	p := s.head
	for p != nil{
		buffer.WriteString(fmt.Sprintf("%v", p.Value))
		p = p.next
	}
	return buffer.String()
}
