package binary_tree

import (
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

func TestBinaryTree(t *testing.T) {
	tree1 := initBinaryTree()
	a := tree1.root
	node2 := a.GetLChild()
	node9 := a.GetRChild().GetLChild().GetRChild()

	fmt.Println("结点2是叶子结点吗? ", node2.IsLeaf())
	fmt.Println("结点9是叶子结点吗? ", node9.IsLeaf())

	fmt.Println("这棵树的结点总数：", tree1.GetSize())

	l := tree1.InOrder()//中序遍历
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*BinTreeNode)
		//fmt.Printf("data:%v\n", *obj)
		fmt.Println(obj)
	}
	result := tree1.Find(6)
	fmt.Printf("结点6的父节点数据：%v\t结点6的右孩子节点数据: %v\n", result.GetParent().GetData(), result.GetRChild().GetData())
}

func TestBinaryTree_PreOrder(t *testing.T) {
	tree := initBinaryTree()
	l := tree.PreOrder()
	l2 := tree.PreOrderByStack()
	assert.Equal(t, l, l2)

	//for e := l.Front(); e != nil; e = e.Next() {
	//	obj, _ := e.Value.(*BinTreeNode)
	//	fmt.Println(obj)
	//}
}

func TestBinaryTree_InOrder(t *testing.T) {
	tree := initBinaryTree()
	l := tree.InOrder()
	l2 := tree.InOrderByStack()
	assert.Equal(t, l, l2)

	//for e := l.Front(); e != nil; e = e.Next() {
	//	obj, _ := e.Value.(*BinTreeNode)
	//	fmt.Println(obj)
	//}
}

func TestBinaryTree_PostOrder(t *testing.T) {
	tree := initBinaryTree()
	l := tree.PostOrder()
	l2 := tree.PostOrderByStack()
	assert.Equal(t, l, l2)

	//for e := l.Front(); e != nil; e = e.Next() {
	//	obj, _ := e.Value.(*BinTreeNode)
	//	fmt.Println(obj)
	//}
}
