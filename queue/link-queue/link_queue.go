package link_queue

import (
	"bytes"
	"fmt"
)

//用双向链表实现队列

// 数据点
type Item interface {

}

type Queue interface {
	Enqueue(t Item)
	Dequeue() Item
	Pop() Item
	Size() int
	IsEmpty() bool
}

// 定义每个节点
type Node struct {
	Value Item
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
func (queue *LinkedQueue) IsEmpty() bool{
	return queue.head == nil
}

// pop
func (queue *LinkedQueue) Pop() Item{
	if queue.IsEmpty(){
		panic("queue is nil")
	}
	return queue.head.Value
}

// Enqueue
func (queue *LinkedQueue) Enqueue(value Item)  {
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
func (queue *LinkedQueue) Dequeue() Item{
	if queue.IsEmpty(){
		panic("queue is nil")
	}
	firstNode := queue.head
	val := firstNode.Value
	queue.head = firstNode.next
	firstNode.next = nil
	firstNode.Value = nil
	queue.size --
	firstNode = nil
	return val
}

//新建队列，即创建Node指针head，用来指向链表第一个结点，初始为空
func NewLinkQueue() *LinkedQueue{
	return &LinkedQueue{
		head: nil, //head指向头部结点
		tail: nil, //tail指向尾部结点
		size: 0,
	}
}

// 遍历显示队列中所有元素
func (queue *LinkedQueue)String() string{
	if queue.IsEmpty(){
		return ""
	}
	var buffer bytes.Buffer
	p := queue.head
	for p != nil{
		buffer.WriteString(fmt.Sprintf("%v", p.Value))
		p = p.next
	}
	return buffer.String()
}
