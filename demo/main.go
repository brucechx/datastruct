package main

import "fmt"

type Node struct {
	val int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func demo(head *Node) *Node {
	// list reverse
	var pre *Node
	curr := head
	for curr != nil{
		pre, curr, curr.next = curr, curr.next, pre
	}
	return pre
}

func main() {
	l := List{}
	l.head = &Node{val:1}
	l.head.next = &Node{val:2}
	newHead := demo(l.head)
	for ;newHead!=nil;newHead=newHead.next{
		fmt.Println(newHead.val)
	}
}
