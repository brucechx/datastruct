package single_list

import (
	"bytes"
	"fmt"
)

type Queue interface {
	IsEmpty() bool
	Length() int
	Append(data interface{})
	InsertHead(data interface{})
	Delete(data interface{})
	Traverse()
	Reverse()
	Exist(data interface{}) bool
}

// 链表结点
type Node struct {
	data interface{} // 数据
	next *Node // 下一个结点
}

// 链表
type LinkList struct {
	head *Node // 头结点
	tail *Node // 尾结点
	size int // 大小
}

//新建空链表，即创建Node指针head，用来指向链表第一个结点，初始为空
func NewLinkList() *LinkList{
	return &LinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

//是否为空链表
func (l *LinkList) IsEmpty() bool{
	return l.size == 0
}

//链表长度
func (l *LinkList) Length() int{
	return l.size
}

//在链表尾部添加数据
func(l *LinkList) Append(data interface{}){
	n := &Node{
		data: data,
		next: nil,
	}
	if l.IsEmpty(){
		l.head = n
		l.tail = n
	}else {
		l.tail.next = n
		l.tail = n
	}
	l.size++
}

//在链表头部插入数据
func(l *LinkList) InsertHead(data interface{}){
	n := &Node{
		data: data,
		next: l.head,
	}
	l.head = n
	if l.IsEmpty(){
		l.tail = n
	}
	l.size++
}

//在指定结点后面插入数据
func(l *LinkList) insertAfterNode(n *Node, data interface{}) error{
	if ! l.existNode(n){
		return fmt.Errorf("%v node does not exit in list", n.data)
	}
	newNode := &Node{
		data: data,
		next: nil,
	}
	nNext := n.next
	n.next = newNode
	newNode.next = nNext
	l.size++
	return nil
}

//在第一次出现指定数据的结点后插入数据,若链表中无该数据，返回false
func(l *LinkList) InsertAfterData(data interface{}, newData interface{}) error{
	n, err := l.getFirstNode(data)
	if err != nil{
		return err
	}
	newNode := &Node{
		data: newData,
		next: nil,
	}
	nNext := n.next
	n.next = newNode
	newNode.next = nNext
	l.size++
	return nil
}

//在指定下标处插入数据
func(l *LinkList) InsertNodeByPosition(position int, data interface{}) error{
	n, err := l.getNodeByPosition(position)
	if err != nil{
		return err
	}
	return l.insertAfterNode(n, data)
}

//删除指定结点
func(l *LinkList) deleteNode(n *Node) error{
	if ! l.existNode(n){
		return fmt.Errorf("%v node does not exit in list", n.data)
	}

	if n == l.head{ // 如果是头结点
		l.head = l.head.next
		l.size--
		return nil
	}
	if n == l.tail{ // 如果是尾结点
		//寻找到尾结点前一个结点
		p := l.head
		for p.next != l.tail{
			p = p.next
		}
		p.next = nil
		l.tail = p
		l.size--
		return nil
	}
	// 是中间结点; 找到n的前一个结点
	p := l.head
	for p.next != n{
		p = p.next
	}
	p.next = n.next
	n = nil
	l.size--
	return nil
}

//删除第一个含指定数据的结点
func(l *LinkList) Delete(data interface{}) error{
	if n, err := l.getFirstNode(data); err != nil{
		return err
	}else {
		return l.deleteNode(n)
	}
}

// 是否含有指定数据点
func (l *LinkList) Exist(data interface{}) bool{
	for p:=l.head; p != nil; p = p.next{
		if p.data == data{
			return true
		}
	}
	return false
}

// 是否含有指定结点点
func (l *LinkList) existNode(n *Node) bool{
	for p:=l.head; p != nil; p = p.next{
		if p == n{
			return true
		}
	}
	return false
}

//获取含有指定数据的第一个结点
func (l *LinkList) getFirstNode(data interface{}) (*Node, error){
	for p:=l.head; p != nil; p = p.next{
		if p.data == data{
			return p, nil
		}
	}
	return nil, fmt.Errorf("%v node does not exit in list", data)
}

func (l *LinkList) getNodeByPosition(position int) (*Node, error){
	if position > l.size{
		return nil, fmt.Errorf("position:%d is larger than list size:%d", position, l.size)
	}
	p := l.head
	for i:=0; i<position; i++{
		p = p.next
	}
	return p, nil
}

//遍历链表
func(l *LinkList) Traverse(){
	fmt.Println(l.String())
}

//单向链表反转
func (l *LinkList) Reverse(){
	if l.head == nil || l.head.next == nil{
		return
	}
	curr := l.head
	var pre *Node
	for curr != nil{
		pre, curr, curr.next = curr, curr.next, pre
	}
	l.head = pre
}

func (l *LinkList) reverse2(){
	if l.head == nil || l.head.next == nil{
		return
	}
	var pre, curr, tmp *Node // tmp 作为零时结点
	pre = l.head
	curr = pre.next
	for curr != nil{
		tmp = curr.next // 当前结点的下一个结点作为零时结点
		curr.next = pre // 当前节点的前一个节点作为当前节点的下一个节点，刚好实现反转
		pre = curr // 此时当前节点则为下一次反转的当前节点的前一个节点
		curr = tmp // 当前节点的下一个节点，此时为下一次反转的当前节点
	}
	l.head.next = nil
	l.head = pre // 当前节点的前一个节点作为头节点，即最开始的节点
}

func (l *LinkList) reverse3(){
	var pre *Node
	var next *Node
	curr := l.head
	for curr != nil{
		next = curr.next
		curr.next = pre
		pre = curr
		curr = next
	}
	l.head = pre
}


func (l *LinkList) String() string{
	var buffer bytes.Buffer
	for p:=l.head; p != nil; p = p.next{
		buffer.WriteString(fmt.Sprintf("%v", p.data))
	}
	return buffer.String()
}



