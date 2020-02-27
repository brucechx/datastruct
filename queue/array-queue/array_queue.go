package array_queue

import (
	"bytes"
	"fmt"
)

// 用切片实现队列

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

type ArrayQueue struct {
	Items []Item
}

// 创建队列
func NewQueue() *ArrayQueue {
	return &ArrayQueue{
		Items: []Item{},
	}
}

// Enqueue
func (q *ArrayQueue)Enqueue(t Item)  {
	q.Items = append(q.Items, t)
}

// pop
func (q *ArrayQueue)Pop() Item {
	if q.IsEmpty(){
		panic("queue is empty")
	}
	node := q.Items[0]
	return node
}

// Dequeue
func (q *ArrayQueue)Dequeue() Item {
	if q.IsEmpty(){
		panic("queue is empty")
	}
	node := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return node
}

// isEmpty
func (q *ArrayQueue)IsEmpty() bool {
	return len(q.Items) == 0
}

// Size
func (q *ArrayQueue)Size() int {
	return len(q.Items)
}

// 遍历显示队列中的所有元素
func (q *ArrayQueue)String() string{
	var buffer bytes.Buffer
	for _, v := range q.Items{
		buffer.WriteString(fmt.Sprintf("%v", v))
	}
	return buffer.String()
}
