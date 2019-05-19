package simp_avl

import (
	"sync"
	"container/list"
)

// 平衡树
type AvlTree struct {
	root *Node
	lock sync.RWMutex
}

// New simply returns a new (empty) tree
func New() *AvlTree {
	return &AvlTree{}
}

// Insert inserts the Item t in the tree
func (avl *AvlTree) Insert(key int, value interface{}) {
	avl.lock.Lock()
	defer avl.lock.Unlock()
	n := &Node{key, value, nil, nil, 0}
	avl.root = insertNode(avl.root, n)
}

func insertNode(root, newNode *Node) *Node{
	if root == nil{
		root = newNode
		return root
	}
	// 左
	if newNode.key < root.key{
		root.left = insertNode(root.left, newNode)
		if root.left.getHeight() - root.right.getHeight() == 2{
			if newNode.key < root.left.key{  // 左左
				root = rotateWithLeftChild(root)
			}else {
				root = doubleWithLeftChild(root)
			}
		}

	}else if newNode.key > root.key {  // 右
		root.right = insertNode(root.right, newNode)
		if root.right.getHeight() - root.left.getHeight() == 2{
			if newNode.key > root.right.key{
				root = rotateWithRightChild(root)
			}else {
				root = doubleWithRightChild(root)
			}
		}

	}
	root.height = max(root.left.getHeight(), root.right.getHeight()) + 1
	return root
}

// 左左 -- 旋转
/*
		  k2                          k1
		 /  \                        /  \
		k1   z   ==>                x    k2
	   /  \                        /    /  \
	  x    y                      X    y    z
	 /
    X
*/
//  左左： 我们把k1变成这棵树的根节点，因为k2大于k1，把k2置于k1的右子树上，而原本在k1右子树的Y大于k1，小于k2，就把Y置于k2的左子树上

func rotateWithLeftChild(k2 *Node) *Node{
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	k2.height = max(k2.left.getHeight(), k2.right.getHeight()) + 1
	k1.height = max(k1.left.getHeight(), k2.getHeight()) + 1
	return k1
}

// 左右 -- 旋转
/*
		   k3                          k3                       k2
		 /    \    右右                /  \       左左          /   \
		k1(r)  d   ==>               k2    d     ====>       k1    k3
	   /  \                         / \                      / \   / \
	  a    k2                      k1  c                    a   b c   d
           / \                    / \
		  b	  c                  a  b
*/
func doubleWithLeftChild(k3 *Node) *Node{
	k3.left = rotateWithRightChild(k3.left)
	return rotateWithLeftChild(k3)
}

func rotateWithRightChild(k1 *Node) *Node{
	k2 := k1.right
	k1.right = k2.left
	k2.left = k1
	k1.height = max(k1.left.getHeight(), k1.right.getHeight() ) + 1
	k2.height = max(k1.getHeight(), k2.getHeight()) + 1
	return k2
}

func doubleWithRightChild(k1 *Node) *Node{
	k1.right = rotateWithLeftChild(k1.right)
	return rotateWithRightChild(k1)
}


// max gives returns the maximum of a and b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//先序遍历，并保存在链表里
func (avl *AvlTree) PreOrder() *list.List {
	traversal := list.New()
	preOrder(avl.root, traversal)
	return traversal
}

func preOrder(rt *Node, l *list.List) {
	if rt == nil {
		return
	}
	l.PushBack(rt)
	preOrder(rt.left, l)
	preOrder(rt.right, l)
}

//中序遍历，并保存在链表里
func (avl *AvlTree) InOrder() *list.List {
	traversal := list.New()
	inOrder(avl.root, traversal)
	return traversal
}

func inOrder(rt *Node, l *list.List) {
	if rt == nil {
		return
	}
	inOrder(rt.left, l)
	l.PushBack(rt)
	inOrder(rt.right, l)
}



//按层遍历，并保存在链表里
func (avl *AvlTree) SeqTraverse() *list.List {
	traversal := list.New()
	seqTraverse(avl.root, traversal)
	return traversal
}

func seqTraverse(rt *Node, l *list.List) {
	if rt == nil {
		return
	}
	currentSlice := []*Node{rt}
	for len(currentSlice) != 0{
		var nextSlice []*Node
		for _, val := range currentSlice{
			if val.left != nil{
				nextSlice = append(nextSlice, val.left)
			}
			if val.right != nil{
				nextSlice = append(nextSlice, val.right)
			}
			l.PushBack(val)
		}
		currentSlice = nextSlice
	}
	return
}


