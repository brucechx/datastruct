package main

import "fmt"

// 链表实现
// 先入后出

type Node struct {
	Value interface{}
	pre *Node
	next *Node
}

type Stack struct {
	head *Node
	tail *Node
	size int
}
//新建队列，即创建Node指针head，用来指向链表第一个结点，初始为空
func CreateLinkStack() Stack{
	l:=Stack{}
	l.head=nil //head指向头部结点
	l.tail=nil //tail指向尾部结点
	l.size=0
	return l
}

// Push(e Element)    //向栈顶添加元素,加在头前
func (s *Stack) Push(v interface{}){
	newNode := &Node{
		Value:v,
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
func (s *Stack) IsEmpty() bool{
	return s.size == 0
}

//	Pop()   Element    //移除栈顶元素
func (s *Stack)Pop()  interface{}{
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
func (s *Stack) Top() interface{}{
	if s.IsEmpty(){
		panic("栈为空")
	}
	return s.head.Value
}

//	Clear()  bool       //清空栈
func (s *Stack)clear()  {
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
func (s *Stack)Size() int  {
	return s.size
}

// 遍历显示队列中所有元素
func (s *Stack)printInfo()  {
	if s.IsEmpty(){
		fmt.Println("stack is empty")
		return
	}
	p := s.head
	for p != nil{
		fmt.Print(p.Value)
		p = p.next
	}
	fmt.Println()

}
func main() {
	s := CreateLinkStack()
	fmt.Println("栈创建成功：")
	fmt.Println("size: ", s.size)
	s.printInfo()

	fmt.Println("---------------------------现在开始测试插入数据-------------------------")
	fmt.Println("栈中加一个int 1")
	s.Push(1)
	fmt.Println("size: ", s.Size())
	s.printInfo()
	fmt.Println("栈中加一个int 2")
	s.Push(2)
	fmt.Println("size: ", s.Size())
	s.printInfo()

	fmt.Println("---------------------------现在开始测试出栈数据-------------------------")
	fmt.Println("栈头数据")
	fmt.Println(s.Top())
	fmt.Println("size: ", s.Size())
	s.printInfo()

	fmt.Println("出栈")
	fmt.Println(s.Pop())
	fmt.Println("size: ", s.Size())
	s.printInfo()

	fmt.Println("---------------------------现在开始测试清空栈数据-------------------------")
	s.clear()
	fmt.Println("size: ", s.Size())
	s.printInfo()
}
// https://blog.csdn.net/skh2015java/article/details/79231727
