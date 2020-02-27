package binary_tree

import (
	"bytes"
	"container/list"
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

/*
	 		 1
          /      \
		2         5
	     \       /  \
          3     6    7
         /       \    \
        4         9    8
 */

func initBinaryTree() *binaryTree{
	a := NewBinTreeNode(1)
	tree1 := NewBinaryTree(a)
	a.SetLChild(NewBinTreeNode(2))
	a.SetRChild(NewBinTreeNode(5))
	a.GetLChild().SetRChild(NewBinTreeNode(3))
	a.GetLChild().GetRChild().SetLChild(NewBinTreeNode(4))
	a.GetRChild().SetLChild(NewBinTreeNode(6))
	a.GetRChild().SetRChild(NewBinTreeNode(7))
	a.GetRChild().GetLChild().SetRChild(NewBinTreeNode(9))
	a.GetRChild().GetRChild().SetRChild(NewBinTreeNode(8))
	return tree1
}

func listString(l *list.List) string{
	var buffer bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*BinTreeNode)
		buffer.WriteString(fmt.Sprintf("%v", obj.data))
	}
	res := buffer.String()
	return res
}

func TestBinaryTree(t *testing.T) {
	tree1 := initBinaryTree()
	a := tree1.root
	node2 := a.GetLChild()
	node9 := a.GetRChild().GetLChild().GetRChild()
	assert.Equal(t, false, node2.IsLeaf())
	assert.Equal(t, true, node9.IsLeaf())
	assert.Equal(t,9, tree1.GetSize()) //这棵树的结点总数
	result := tree1.Find(6)
	assert.Equal(t, 5, result.GetParent().GetData()) // 结点6的父节点数据
	assert.Equal(t, 9, result.GetRChild().GetData()) // 结点6的右孩子节点数据
}

func TestBinaryTree_PreOrder(t *testing.T) {
	tree := initBinaryTree()
	l := tree.PreOrder()
	l2 := tree.PreOrderByStack()
	assert.Equal(t, l, l2)
	assert.Equal(t, "123456978", listString(l2))
}

func TestBinaryTree_InOrder(t *testing.T) {
	tree := initBinaryTree()
	l := tree.InOrder()
	l2 := tree.InOrderByStack()
	assert.Equal(t, l, l2)
	assert.Equal(t, "243169578", listString(l2))
}

func TestBinaryTree_PostOrder(t *testing.T) {
	tree := initBinaryTree()
	l := tree.PostOrder()
	l2 := tree.PostOrderByStack()
	assert.Equal(t, l, l2)
	assert.Equal(t, "432968751", listString(l2))
}

func TestBinaryTree_SeqTraverse(t *testing.T) {
	tree := initBinaryTree()
	l := tree.SeqTraverse()
	assert.Equal(t, "125367498", listString(l))
}