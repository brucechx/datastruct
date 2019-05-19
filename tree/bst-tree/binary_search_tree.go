package bst_tree

import (
	"sync"
	"fmt"
)

// Node a single node that composes the tree
type Node struct {
	key   int  // 中序遍历的节点序号
	value interface{} // 节点存储的值
	left  *Node //left
	right *Node //right
}

// ItemBinarySearchTree the binary search tree of Items
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

// Insert inserts the Item t in the tree
func (bst *ItemBinarySearchTree) Insert(key int, value interface{}) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// InOrderTraverse visits all nodes with in-order traversing
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(interface{})) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

// internal recursive function to traverse in order
func inOrderTraverse(n *Node, f func(interface{})) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse visits all nodes with pre-order traversing
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

// internal recursive function to traverse pre order
func preOrderTraverse(n *Node, f func(interface{})) {
	if n != nil {
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}


// PostOrderTraverse visits all nodes with post-order traversing
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

// internal recursive function to traverse post order
func postOrderTraverse(n *Node, f func(interface{})) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}

// Min returns the Item with min value stored in the tree
func (bst *ItemBinarySearchTree) Min() *interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// Max returns the Item with max value stored in the tree
func (bst *ItemBinarySearchTree) Max() *interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

// Search returns true if the Item t exists in the tree
func (bst *ItemBinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

// internal recursive function to search an item in the tree
func search(n *Node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return search(n.left, key)
	}
	if key > n.key {
		return search(n.right, key)
	}
	return true
}

// Remove removes the Item with key `key` from the tree
func (bst *ItemBinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

// 递归删除节点
// internal recursive function to remove an item
func remove(node *Node, key int) *Node {
	// 要删除的节点不存在
	if node == nil {
		return nil
	}
	// 寻找节点
	// 要删除的节点在左侧
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	// 要删除的节点在右侧
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}
	// 判断节点类型
	// 要删除的节点是叶子节点，直接删除
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	// 要删除的节点只有一个节点，删除自身
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}

	// 要删除的节点有 2 个子节点，找到右子树的最左节点，替换当前节点
	leftmostrightside := node.right
	for {
		// 一直遍历找到最左节点
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	// 使用右子树的最左节点替换当前节点，即删除当前节点
	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = remove(node.right, node.key)
	return node
}


// String prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}


// ===========================================================

// 迭代查找
func (bst *ItemBinarySearchTree) IterativeSearch(key int) bool{
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return iterativeSearch(bst.root, key)
}

func iterativeSearch(n *Node, key int) bool{
	for n != nil{
		if key == n.key{
			return true
		}
		if key < n.key{
			n = n.left
		}else {
			n = n.right
		}
	}
	return false
}


func (bst *ItemBinarySearchTree)SearchNode(key int) *Node{
	n := bst.root
	for n != nil{
		if key == n.key{
			return n
		}
		if key < n.key{
			n = n.left
		}else {
			n = n.right
		}
	}
	return nil
}

// getParentNode
func (bst *ItemBinarySearchTree) GetParentNode(key int) *Node{
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return bst.getParentNode(bst.root, key)
}

func (bst *ItemBinarySearchTree) getParentNode(n *Node, key int) *Node{
	if n.left == nil && n.right == nil{
		return nil
	}
	for n != nil{
		if n.left != nil && n.left.key == key{
			return n
		}
		if n.right != nil && n.right.key == key{
			return n
		}
		if key < n.key{
			n = n.left
		}
		if key > n.key{
			n = n.right
		}
		if key == n.key{
			return nil
		}
	}
	return nil
}

// 前驱结点
func (bst *ItemBinarySearchTree) predecessor(key int) *Node {
	x := bst.SearchNode(key)
	if x.left != nil{
		return bst.MaxiMum(x.left.key)
	}
	y := bst.GetParentNode(key)
	for y != nil && x == y.left{
		x = y
		y = bst.GetParentNode(y.key)
	}
	return y
}

// 后继结点
func (bst *ItemBinarySearchTree) Successor(key int) *Node{
	x := bst.SearchNode(key)
	if x.right != nil{
		return bst.MiniMum(x.right.key)
	}
	y := bst.GetParentNode(key)
	for y != nil && x == y.right{
		x = y
		y = bst.GetParentNode(y.key)
	}
	return y
}

func (bst *ItemBinarySearchTree) MiniMum(key int) *Node {
	n := bst.SearchNode(key)
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return n
		}
		n = n.left
	}
}

func (bst *ItemBinarySearchTree) MaxiMum(key int) *Node {
	n := bst.SearchNode(key)
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return n
		}
		n = n.right
	}
}

