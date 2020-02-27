package bst_tree

import (
	"fmt"
	"strings"
)

// 二叉搜索数
/*
// 前驱结点
1.1. 若一个节点有左子树，那么该节点的前驱节点是其左子树中val值最大的节点（也就是左子树中所谓的rightMostNode）
1.2. 若一个节点没有左子树，那么判断该节点和其父节点的关系
	1.2.1 若该节点是其父节点的右边孩子，那么该节点的前驱结点即为其父节点。
	1.2.2 若该节点是其父节点的左边孩子，那么需要沿着其父亲节点一直向树的顶端寻找，直到找到一个节点P，P节点是其父节点Q的右边孩子，那么Q就是该节点的前驱节点

// 后继结点
2.1. 若一个节点有右子树，那么该节点的后继节点是其右子树中val值最小的节点（也就是右子树中所谓的leftMostNode）
2.2. 若一个节点没有右子树，那么判断该节点和其父节点的关系
	2.2.1 若该节点是其父节点的左边孩子，那么该节点的后继结点即为其父节点
	2.2.2 若该节点是其父节点的右边孩子，那么需要沿着其父亲节点一直向树的顶端寻找，直到找到一个节点P，P节点是其父节点Q的左边孩子，那么Q就是该节点的后继节点

// 删除
二叉搜索树的结点删除比插入较为复杂，总体来说，结点的删除可归结为三种情况：
3.1. 如果结点z没有孩子节点，那么只需简单地将其删除，并修改父节点，用NIL来替换z；
3.2. 如果结点z只有一个孩子，那么将这个孩子节点提升到z的位置，并修改z的父节点，用z的孩子替换z；
3.3. 如果结点z有2个孩子，那么查找z的后继y，此外后继一定在z的右子树中，然后让y替换z。
	3.3.1 如果y是z的右孩子，那么用y替换z, 并仅留下y的右孩子
	3.3.2 如果y不是z的右孩子，那么先用y的右孩子替换y, 再用y替换z。
*/

// 二叉树对外接口
type BstT interface {
	Insert(key int, i Item) // 加入结点
	Search(key int) (Item, error) // 查找结点
	Exist(key int) bool // 结点是否存在
	Remove(key int) error // 输出结点
	MaxNode() Item // 树中最大结点
	MinNode() Item // 树中最小结点
	PreOrder(callback func (i Item)) // 前序遍历
	MidOrder(callback func (i Item)) // 中续遍历
	PostOrder(callback func (i Item)) // 后续遍历
	PreSuccessor(key int) (Item, error) // 前驱结点
	Successor(key int) (Item, error) // 后继结点
}

// 数据结点
type Item interface {

}

// 树结点
type Node struct {
	key int // 结点序号
	val Item // 结点数据
	left *Node // 左结点
	right *Node // 右结点
}

// 二叉搜索树
type BstTree struct {
	root *Node // 根结点
}

func NewBstTree(key int, val Item) *BstTree{
	return &BstTree{
		root: &Node{
			key:   key,
			val:   val,
			left:  nil,
			right: nil,
		},
	}
}

func (b *BstTree)Insert(key int, i Item){
	newNode := &Node{
		key:   key,
		val:   i,
		left:  nil,
		right: nil,
	}
	if b.root == nil{
		b.root = newNode
		return
	}
	b.insert(b.root, newNode)
}

func (b *BstTree)insert(curr, new *Node){
	if new.key < curr.key{
		if curr.left == nil{
			curr.left = new
		}else {
			b.insert(curr.left, new)
		}
	}else {
		if curr.right == nil{
			curr.right = new
		}else {
			b.insert(curr.right, new)
		}
	}
}

func (b *BstTree)Search(key int) (Item, error){
	if node, err :=  b.search(b.root, key); err != nil{
		return nil, err
	}else {
		return node.val, nil
	}
}

func (b *BstTree)search(n *Node, key int) (*Node, error){
	for n != nil{
		if key == n.key{
			return n, nil
		}
		if key < n.key{
			n = n.left
		}else {
			n = n.right
		}
	}
	return nil, fmt.Errorf("node:%d does not exist", key)
}


func (b *BstTree)Exist(key int) bool{
	if _, err := b.search(b.root, key); err != nil{
		return false
	}else {
		return true
	}
}


func (b *BstTree) MaxNode() Item{
	node := b.maxNode(b.root)
	if node != nil{
		return node.val
	}
	return nil
}

func (b *BstTree) maxNode(n *Node) *Node{
	if n == nil{
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (b *BstTree) MinNode() Item{
	node := b.minNode(b.root)
	if node == nil{
		return nil
	}
	return node.val
}

func (b *BstTree) minNode(n *Node) *Node{
	for n.left != nil{
		n = n.left
	}
	return n
}

func (b *BstTree)PreOrder(callback func (i Item))  {
	b.preOrder(b.root, callback)
}

func (b *BstTree)preOrder(n *Node, callback func(i Item)){
	if n != nil{
		callback(n.val)
		b.preOrder(n.left, callback)
		b.preOrder(n.right, callback)
	}
}

func (b *BstTree)MidOrder(callback func (i Item))  {
	b.midOrder(b.root, callback)
}

func (b *BstTree)midOrder(n *Node, callback func(i Item)){
	if n != nil{
		b.midOrder(n.left, callback)
		callback(n.val)
		b.midOrder(n.right, callback)
	}
}

func (b *BstTree)PostOrder(callback func (i Item))  {
	b.postOrder(b.root, callback)
}

func (b *BstTree)postOrder(n *Node, callback func(i Item)){
	if n != nil{
		b.postOrder(n.left, callback)
		b.postOrder(n.right, callback)
		callback(n.val)
	}
}

func (b *BstTree)ParentNode(key int) *Node{
	return b.parentNode(b.root, key)
}

func (b *BstTree)parentNode(n *Node, key int) *Node{
	if n.left == nil && n.right == nil{
		return nil
	}
	for n != nil{
		if n.key == key{
			return nil
		}
		if n.left != nil && n.left.key == key{
			return n
		}
		if n.right != nil && n.right.key == key{
			return n
		}
		if n.left != nil && key < n.key{
			n = n.left
			continue
		}
		if n.right != nil && key > n.key{
			n = n.right
			continue
		}
	}
	return nil
}

func (b *BstTree)PreSuccessor(key int) (Item, error) {
	node, err := b.preSuccessor(key)
	if err != nil || node == nil{
		return node, err
	}
	return node.val, nil
}

// 前驱结点
func (b *BstTree)preSuccessor(key int) (*Node, error){
	currentNode, err := b.search(b.root, key)
	if err != nil{
		return nil, err
	}
	// 1.1
	if currentNode.left != nil{
		return b.maxNode(currentNode.left), nil
	}
	// 1.2
	parentNode := b.ParentNode(key)
	if parentNode == nil{
		return nil, nil
	}
	// 1.2.1
	if currentNode == parentNode.right{
		return parentNode, nil
	}
	// 1.2.2
	for parentNode != nil && currentNode == parentNode.left{
		currentNode = parentNode
		parentNode = b.ParentNode(parentNode.key)
	}
	return parentNode, nil
}

// 后继结点
func (b *BstTree)Successor(key int) (Item, error) {
	node, err := b.successor(key)
	if err != nil || node == nil{
		return nil, err
	}
	return node.val, nil
}

func (b *BstTree)successor(key int) (*Node, error){
	currentNode, err := b.search(b.root, key)
	if err != nil{
		return nil, err
	}
	// 2.1
	if currentNode.right != nil{
		return b.minNode(currentNode.right), nil
	}
	// 2.2
	parentNode := b.ParentNode(key)
	if parentNode == nil{
		return nil, nil
	}
	// 2.2.1
	if currentNode == parentNode.left{
		return parentNode, nil
	}
	// 2.2.2
	for parentNode != nil && currentNode == parentNode.right{
		currentNode = parentNode
		parentNode = b.ParentNode(parentNode.key)
	}
	return parentNode, nil
}

func (b *BstTree)Remove(key int) error{
	currentNode, err := b.search(b.root, key)
	if err != nil{
		return err
	}
	// 3.1
	if currentNode.left == nil && currentNode.right == nil{
		parent := b.ParentNode(key)
		if currentNode == parent.left{
			parent.left = nil
		}else {
			parent.right = nil
		}
		currentNode = nil
		return nil
	}
	// 3.2
	parent := b.ParentNode(key)
	if currentNode.left == nil && currentNode.right != nil{
		if currentNode == parent.left{
			parent.left = currentNode.right
		}else {
			parent.right = currentNode.right
		}
		currentNode = nil
		return nil
	}
	if currentNode.right == nil && currentNode.left != nil{
		if currentNode == parent.left{
			parent.left = currentNode.left
		}else {
			parent.right = currentNode.left
		}
		currentNode = nil
		return nil
	}
	// 3.3
	successor, err := b.successor(key)
	if err != nil{
		return err
	}
	// 3.3.1
	if successor == currentNode.right{
		if currentNode == parent.left{
			parent.left = successor
		}else {
			parent.right = successor
		}
		successor.left = currentNode.left
		return nil
	}
	// 3.3.2
	// 那么先用y的右孩子替换y, 再用y替换z。
	successor = successor.right
	if currentNode == parent.left{
		parent.left = successor
	}else {
		parent.right = successor
	}
	return nil
}


// Remove removes the Item with key `key` from the tree
func (b *BstTree) Remove2(key int) {
	remove(b.root, key)
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
	leftMostRightSide := node.right
	for {
		// 一直遍历找到最左节点
		//find smallest value on the right side
		if leftMostRightSide != nil && leftMostRightSide.left != nil {
			leftMostRightSide = leftMostRightSide.left
		} else {
			break
		}
	}
	// 使用右子树的最左节点替换当前节点，即删除当前节点
	if leftMostRightSide != nil{
		node.key, node.val = leftMostRightSide.key, leftMostRightSide.val
	}
	node.right = remove(node.right, node.key)
	return node
}

func (b *BstTree)String() string{
	var result []string
	b.MidOrder(func(i Item) {
		result = append(result, fmt.Sprintf("%v", i))
	})
	res := strings.Join(result, ",")
	return res
}

// String prints a visual representation of the tree
func (b *BstTree) String2() {
	fmt.Println("------------------------------------------------")
	stringify(b.root, 0)
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