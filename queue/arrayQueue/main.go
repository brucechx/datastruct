package main

import "fmt"

// 用切片实现队列
type Item interface {

}

type ArrayQueue struct {
	Items []Item
}

// 创建队列
func (q *ArrayQueue)CreateQueue() *ArrayQueue {
	q.Items = []Item{}
	return q
}

// Enqueue
func (q *ArrayQueue)Enqueue(t Item)  {
	q.Items = append(q.Items, t)
}

// pop
func (q *ArrayQueue)Pop() Item {
	if q.isEmpty(){
		panic("queue is empty")
	}
	node := q.Items[0]
	return node
}

// Dequeue
func (q *ArrayQueue)Dequeue() Item {
	if q.isEmpty(){
		panic("queue is empty")
	}
	node := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return node
}

// isEmpty
func (q *ArrayQueue)isEmpty() bool {
	return len(q.Items) == 0
}

// Size
func (q *ArrayQueue)Size() int {
	return len(q.Items)
}

// 遍历显示队列中的所有元素
func (q *ArrayQueue)printInfo(){
	for _, v := range q.Items{
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	q := &ArrayQueue{}
	q.CreateQueue()
	fmt.Println("链表创建成功：")
	fmt.Println()
	fmt.Println("---------------------------现在开始测试插入数据-------------------------")
	fmt.Println("队列加一个int 1")
	q.Enqueue(1)
	q.printInfo()

	q.Enqueue(2)
	q.printInfo()

	fmt.Println("---------------------------现在开始测试出队数据-------------------------")
	fmt.Println("出队:")
	fmt.Println(q.Pop())
	fmt.Println("当前队列大小：",q.Size())
	q.printInfo()

	fmt.Println("---------------------------现在开始测试删除数据-------------------------")
	fmt.Println("删除:")
	q.Dequeue()
	fmt.Println("当前队列大小：",q.Size())
	q.printInfo()

}