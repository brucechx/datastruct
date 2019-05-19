package main

import "fmt"

// 用slice 实现
// 先进后出
type Item interface {}

type ArrayStack struct {
	Items []Item
}

func (s *ArrayStack)New() *ArrayStack{
	s.Items = []Item{}
	return s
}

// push
func (s *ArrayStack)Push(i Item)  {
	s.Items = append(s.Items, i)
}

// pop
func (s *ArrayStack)Pop() Item {
	if s.isEmpty(){
		panic("栈为空")
	}
	item := s.Items[len(s.Items) - 1]
	s.Items = s.Items[:len(s.Items) - 1]
	return item
}

// isEmpty
func (s *ArrayStack)isEmpty() bool {
	return len(s.Items) == 0
}

func (s *ArrayStack)Size() int {
	return len(s.Items)
}

// 遍历显示队列中的所有元素
func (q *ArrayStack)printInfo(){
	for _, v := range q.Items{
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	q := &ArrayStack{}
	q.New()
	fmt.Println("栈创建成功：")
	fmt.Println()
	fmt.Println("---------------------------现在开始测试插入数据-------------------------")
	fmt.Println("压栈加一个int 1")
	q.Push(1)
	q.printInfo()

	fmt.Println("压栈加一个int 2")
	q.Push(2)
	q.printInfo()

	fmt.Println("---------------------------现在开始测试出栈数据-------------------------")
	fmt.Println("出栈:")
	fmt.Println(q.Pop())
	fmt.Println("当前栈大小：",q.Size())
	q.printInfo()

}