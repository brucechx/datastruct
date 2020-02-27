package avl_tree

import (
	"bytes"
	"container/list"
	"fmt"
)

/*
avl 树
*/

type Tree interface {
	Exist(key int) bool
	Get(key int) (interface{}, error)
	Size() int
	GetHeight() int
	Insert(key int, val interface{})
	Delete(key int)
}

// avl 树
type AvlTree struct {
	root *AvlNode
}

func NewAvlTree() *AvlTree{
	return &AvlTree{}
}

func (t *AvlTree) Exist(key int) bool{
	return t.root.exist(key)
}

func (t *AvlTree) Get(key int) (interface{}, error){
	return t.root.get(key)
}

func (t *AvlTree) Size() int{
	return 0
}

func (t *AvlTree) GetHeight() int{
	return t.root.getHeight()
}

func (t *AvlTree) Insert(key int, val interface{}){
	t.root = t.root.insert(key, val)
}

func (t *AvlTree) Delete(key int){
	t.root = t.root.Delete(key)
}

// for debug
func (t *AvlTree) String() string{
	// 按层遍历
	return listString(t.SeqTraverse())
}

// 按层遍历
func (t *AvlTree) SeqTraverse() *list.List{
	traversal := list.New()
	seqTraverse(t.root, traversal)
	return traversal
}

func seqTraverse(rt *AvlNode, l *list.List){
	if rt == nil{
		return
	}
	currentNodes := make([]*AvlNode, 0)
	currentNodes = append(currentNodes, rt)
	for len(currentNodes) != 0{
		nextNodes := make([]*AvlNode, 0)
		for _, val := range currentNodes{
			l.PushBack(val.val)
			if val.left != nil{
				nextNodes = append(nextNodes, val.left)
			}
			if val.right != nil{
				nextNodes = append(nextNodes, val.right)
			}
			l.PushBack(",")
		}
		l.PushBack(";")
		currentNodes = nextNodes
	}
	return
}

func listString(l *list.List) string{
	var buffer bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		buffer.WriteString(fmt.Sprintf("%v", e.Value))
	}
	res := buffer.String()
	return res
}

