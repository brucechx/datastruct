package main

import "fmt"

//用双向链表实现队列

// 定义每个节点
type Node struct {
	Value interface{} // 值为任意类型
	pre *Node
	next *Node
}

// 定义链表
type LinkedQueue struct {
	head *Node
	tail *Node
	size int
}

// 获取队列大小
func (queue *LinkedQueue) Size() int{
	return queue.size
}

// isEmpty
func (queue *LinkedQueue) isEmpty() bool{
	return queue.head == nil
}

// peek
func (queue *LinkedQueue) peek() interface{}{
	if queue.isEmpty(){
		panic("queue is nil")
	}
	return queue.head.Value
}

// add
func (queue *LinkedQueue) add(value interface{})  {
	newNode := &Node{
		Value: value,
		pre: queue.tail,
		next:nil,
	}
	if queue.tail == nil{
		queue.head = newNode
		queue.tail = newNode
	}else {
		queue.tail.next = newNode
		queue.tail = newNode
	}
	queue.size++

}

// del
func (queue *LinkedQueue) del()  {
	if queue.isEmpty(){
		panic("queue is nil")
	}
	firstNode := queue.head
	queue.head = firstNode.next
	firstNode.next = nil
	firstNode.Value = nil
	queue.size --
	firstNode = nil
}

//新建队列，即创建Node指针head，用来指向链表第一个结点，初始为空
func CreateLinkQueue() LinkedQueue{
	l:=LinkedQueue{}
	l.head=nil //head指向头部结点
	l.tail=nil //tail指向尾部结点
	l.size=0
	return l
}

// 遍历显示队列中所有元素
func (queue *LinkedQueue)printInfo()  {
	if queue.isEmpty(){
		fmt.Println("queue is empty")
		return
	}
	p := queue.head
	for p != nil{
		fmt.Print(p.Value)
		p = p.next
	}
	fmt.Println()

}

func main() {
	q := CreateLinkQueue()
	fmt.Println("链表创建成功：")
	q.printInfo()
	fmt.Println()
	fmt.Println("---------------------------现在开始测试插入数据-------------------------")
	fmt.Println("队列加一个int 1")
	q.add(1)
	fmt.Println("当前队列大小：",q.size)
	q.printInfo()

	fmt.Println("队列加一个int 2")
	q.add(2)
	fmt.Println("当前队列大小：",q.size)
	q.printInfo()

	fmt.Println("队列加一个string 3")
	q.add("3")
	fmt.Println("当前队列大小：",q.size)
	q.printInfo()

	fmt.Println("---------------------------现在开始测试出队数据-------------------------")
	fmt.Println("出队:")
	fmt.Println(q.peek())
	fmt.Println("当前队列大小：",q.size)
	q.printInfo()

	fmt.Println("---------------------------现在开始测试删除数据-------------------------")
	fmt.Println("删除:")
	q.del()
	fmt.Println("当前队列大小：",q.size)
	q.printInfo()


}
