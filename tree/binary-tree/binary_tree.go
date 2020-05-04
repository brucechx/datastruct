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
func (tree *binaryTree) GetSize() int {
	return tree.root.size
}

//判断二叉树是否为空
func (tree *binaryTree) IsEmpty() bool {
	return tree.root != nil
}

//获得二叉树根节点
func (tree *binaryTree) GetRoot() *BinTreeNode {
	return tree.root
}

//获得二叉树高度，根节点层为0
func (tree *binaryTree) GetHeight() int {
	return tree.root.height
}

//获得第一个与数据e相等的节点
func (tree *binaryTree) Find(e interface{}) *BinTreeNode {
	if tree.root == nil {
		return nil
	}
	p := tree.root
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
func (tree *binaryTree) PreOrder() *list.List {
	traversal := list.New()
	preOrder(tree.root, traversal)
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

//先序遍历，非递归并， 通过栈，保存先序数据在链表里
func (tree *binaryTree) PreOrderByStack() *list.List {
	traversal := list.New()
	preOrderByStack(tree.root, traversal)
	return traversal
}

/*
	前序遍历； 根左右
    备注：直接入列，随后根左入栈再出栈操作右结点
 */
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
func (tree *binaryTree) InOrder() *list.List {
	traversal := list.New()
	inOrder(tree.root, traversal)
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
func (tree *binaryTree) InOrderByStack() *list.List {
	traversal := list.New()
	inOrderByStack(tree.root, traversal)
	return traversal
}

/*
 	中序遍历： 左根右
	备注：先根左入栈，出栈后再入列
    通过栈实现将数据存入队列中：
		1. 根，左 最大深度依次入栈，直至无左结点
        2. 随后出栈； (右)根(左)出，入队列；
			取其右结点，若不为空，重复1；若为空，重复2
 */
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


// 中序列遍历2, 非递归实现
/*
p = root
st = []  # 用列表模拟实现栈的功能
while p is not None or st:
    while p is not None:
        st.append(p)
        p = p.left
    p = st.pop()
    proc(p.val)
    p = p.right
*/
func (tree *binaryTree) MiddleOrder() *list.List{
	p := tree.root
	st := NewStack()
	l := list.New()
	for p != nil ||  ! st.isEmpty(){
		for p != nil{
			st.Push(p)
			p = p.GetLChild()
		}
		p = st.Pop().(*BinTreeNode)
		l.PushBack(p)
		p = p.GetRChild()
	}
	return l
}

//后序遍历，并保存在链表里
func (tree *binaryTree) PostOrder() *list.List {
	traversal := list.New()
	postOrder(tree.root, traversal)
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
func (tree *binaryTree) PostOrderByStack() *list.List {
	traversal := list.New()
	postOrderByStack(tree.root, traversal)
	return traversal
}

/*
	后续遍历: 左右根
	备注：先根右 左入栈，出栈后再入列
	通过栈实现遍历: 则根入栈，然后右结点入栈，最后左结点入栈；
			出栈的时候刚好是左右根
 */
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
			tmp := stack.Pop().(*BinTreeNode)
			rt = tmp.GetLChild()
		}
	}

	for ! outStack.isEmpty(){
		node := outStack.Pop().(*BinTreeNode)
		if node == nil{
			continue
		}
		l.PushBack(node)
	}
}

// 按层遍历
// 递归
func (tree *binaryTree)levelOrder() [][]interface{}{
	res := make([][]interface{}, 0)
	dfs(tree.root, &res, 0)
	return  res
}

func dfs(node *BinTreeNode, res *[][]interface{}, level int){
	if node == nil{
		return
	}
	if len(*res) == level{
		*res = append(*res, []interface{}{})
	}
	(*res)[level] = append((*res)[level], node.data)
	for _, v := range []*BinTreeNode{node.lChild, node.rChild}{
		dfs(v, res, level+1)
	}
}

func (tree *binaryTree) SeqTraverse() *list.List{
	traversal := list.New()
	seqTraverse(tree.root, traversal)
	return traversal
}

func seqTraverse(rt *BinTreeNode, l *list.List){
	if rt == nil{
		return
	}
	currentNodes := make([]*BinTreeNode, 0)
	currentNodes = append(currentNodes, rt)
	for len(currentNodes) != 0{
		nextNodes := make([]*BinTreeNode, 0)
		for _, val := range currentNodes{
			if val.lChild != nil{
				nextNodes = append(nextNodes, val.lChild)
			}
			if val.rChild != nil{
				nextNodes = append(nextNodes, val.rChild)
			}
			l.PushBack(val)
		}
		currentNodes = nextNodes
	}
	return
}
