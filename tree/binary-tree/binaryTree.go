package binary_tree

import (
	"container/list"
)

//二叉树
type binaryTree struct {
	root   *BinTreeNode //根节点
	height int
	size   int
}

func NewBinaryTree(root *BinTreeNode) *binaryTree {
	return &binaryTree{root: root}
}

//获得二叉树总结点数
func (this *binaryTree) GetSize() int {
	return this.root.size
}

//判断二叉树是否为空
func (this *binaryTree) IsEmpty() bool {
	return this.root != nil
}

//获得二叉树根节点
func (this *binaryTree) GetRoot() *BinTreeNode {
	return this.root
}

//获得二叉树高度，根节点层为0
func (this *binaryTree) GetHeight() int {
	return this.root.height
}

//获得第一个与数据e相等的节点
func (this *binaryTree) Find(e interface{}) *BinTreeNode {
	if this.root == nil {
		return nil
	}
	p := this.root
	return isEqual(e, p)
}

func isEqual(e interface{}, node *BinTreeNode) *BinTreeNode {
	if e == node.GetData() {
		return node
	}

	if node.HasLChild() {
		lp := isEqual(e, node.GetLChild())
		if lp != nil {
			return lp
		}
	}

	if node.HasRChild() {
		rp := isEqual(e, node.GetRChild())
		if rp != nil {
			return rp
		}

	}

	return nil
}

//先序遍历，并保存在链表里
func (this *binaryTree) PreOrder() *list.List {
	traversal := list.New()
	preOrder(this.root, traversal)
	return traversal
}

func preOrder(rt *BinTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	l.PushBack(rt)
	preOrder(rt.GetLChild(), l)
	preOrder(rt.GetRChild(), l)
}

//先序遍历，非递归并保存在链表里
func (this *binaryTree) PreOrderByStack() *list.List {
	traversal := list.New()
	preOrderByStack(this.root, traversal)
	return traversal
}
func preOrderByStack(rt *BinTreeNode, l *list.List){
	var stack Stack
	stack.New()

	for rt != nil || !stack.isEmpty(){
		if rt != nil{
			l.PushBack(rt)
			stack.Push(rt)
			rt = rt.GetLChild()
		}else {
			tmp := stack.Pop()
			rt = tmp.(*BinTreeNode).GetRChild()
		}
	}
}

//中序遍历，并保存在链表里
func (this *binaryTree) InOrder() *list.List {
	traversal := list.New()
	inOrder(this.root, traversal)
	return traversal
}

func inOrder(rt *BinTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	inOrder(rt.GetLChild(), l)
	l.PushBack(rt)
	inOrder(rt.GetRChild(), l)
}

//中序遍历，非递归，并保存在链表里
func (this *binaryTree) InOrderByStack() *list.List {
	traversal := list.New()
	inOrderByStack(this.root, traversal)
	return traversal
}

func inOrderByStack(rt *BinTreeNode, l *list.List) {
	var stack Stack
	stack.New()

	for rt != nil || ! stack.isEmpty(){
		if rt != nil{
			stack.Push(rt)
			rt = rt.GetLChild()
		}else {
			tmp := stack.Pop()
			rt = tmp.(*BinTreeNode)
			l.PushBack(rt)
			rt = rt.GetRChild()
		}

	}
}


//后序遍历，并保存在链表里
func (this *binaryTree) PostOrder() *list.List {
	traversal := list.New()
	postOrder(this.root, traversal)
	return traversal
}

func postOrder(rt *BinTreeNode, l *list.List) {
	if rt == nil {
		return
	}
	postOrder(rt.GetLChild(), l)
	postOrder(rt.GetRChild(), l)
	l.PushBack(rt)
}

//后序遍历，非递归，并保存在链表里
func (this *binaryTree) PostOrderByStack() *list.List {
	traversal := list.New()
	postOrderByStack(this.root, traversal)
	return traversal
}

func postOrderByStack(rt *BinTreeNode, l *list.List) {
	var stack, outStack Stack
	stack.New()
	outStack.New()
	for rt != nil || ! stack.isEmpty(){
		if rt != nil{
			stack.Push(rt)
			outStack.Push(rt)
			rt = rt.GetRChild()
		}else {
			node := NewBinTreeNode(nil)
			outStack.Push(node)
			tmp := stack.Pop()
			rt = tmp.(*BinTreeNode)
			rt = rt.GetLChild()
		}
	}

	for ! outStack.isEmpty(){
		tmp := outStack.Pop()
		node := tmp.(*BinTreeNode)
		if node.data == nil{
			continue
		}
		l.PushBack(node)
	}
}

