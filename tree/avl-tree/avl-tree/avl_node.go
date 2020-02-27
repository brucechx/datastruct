package avl_tree

import (
	"errors"
	"fmt"
)

// avl 树结点
type AvlNode struct {
	key int
	val interface{}
	height int
	left *AvlNode
	right *AvlNode
}

func (n *AvlNode) exist(key int) bool{
	_, err := n.get(key)
	if err != nil{
		return false
	}
	return true
}

func (n *AvlNode) get(key int) (val interface{}, err error){
	if n == nil{
		return nil, errors.New("not found")
	}
	if n.key == key{
		return n.val, nil
	}else if key < n.key{
		return n.left.get(key)
	}else {
		return n.right.get(key)
	}
}

func (n *AvlNode) getHeight() int{
	if n == nil{
		return 0
	}
	return n.height
}

// 左旋
func (n *AvlNode) leftRotate() *AvlNode{
	newRootNode := n.right
	n.right = newRootNode.left
	newRootNode.left = n
	// 更新结点高度
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	newRootNode.height = max(newRootNode.left.getHeight(), newRootNode.right.getHeight()) + 1
	return newRootNode
}

// 右旋
func (n *AvlNode) rightRotate() *AvlNode{
	newRootNode := n.left
	n.left = newRootNode.right
	newRootNode.right = n
	// 更新高度
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	newRootNode.height = max(newRootNode.left.getHeight(), newRootNode.right.getHeight()) + 1
	return newRootNode
}

// 先左再右旋
func (n *AvlNode) leftThenRightRotate() *AvlNode{
	// 1. 以失衡点左结点先左旋转
	sonRootNode := n.left.leftRotate()
	n.left = sonRootNode
	// 2. 再以失衡点右旋转
	return n.rightRotate()
}

// 先右再左旋
func (n *AvlNode) rightThenLeftRotate() *AvlNode{
	// 1. 以失衡点右结点先右旋转
	sonRootNode := n.right.rightRotate()
	n.right = sonRootNode
	// 2. 再以失衡点左旋转
	return n.leftRotate()
}

// 均衡，适配规则
func (n *AvlNode) balance() *AvlNode{
	if n.right.getHeight() - n.left.getHeight() == 2{
		if n.right.right.getHeight() > n.right.left.getHeight(){
			return n.leftRotate()
		}else {
			return n.rightThenLeftRotate()
		}
	}else if n.left.getHeight() - n.right.getHeight() == 2{
		if n.left.left.getHeight() - n.left.right.getHeight() == 2{
			return n.rightRotate()
		}else {
			return n.leftThenRightRotate()
		}
	}
	return n
}

func (n *AvlNode) insert(key int, val interface{}) *AvlNode{
	if n == nil{
		return &AvlNode{
			key:    key,
			val:    val,
			height: 0,
			left:   nil,
			right:  nil,
		}
	}
	if key < n.key{
		n.left = n.left.insert(key, val)
		n = n.balance()
	}else if key > n.key{
		n.right = n.right.insert(key, val)
		n = n.balance()
	}else {
		fmt.Println("node has been exist")
	}
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	return n
}

/*删除元素
*1、如果被删除结点只有一个子结点，就直接将A的子结点连至A的父结点上，并将A删除
*2、如果被删除结点有两个子结点，将该结点右子数内的最小结点取代A。
*3、查看是否平衡,该调整调整
*/
func (n *AvlNode) Delete(key int) *AvlNode{
	if n == nil{
		return n
	}
	if key < n.key{
		n.left = n.left.Delete(key)
	}else if key > n.key{
		n.right = n.right.Delete(key)
	}else { // 找到结点，并删除
		if n.left != nil && n.right != nil{
			minNode := n.right.getMinNode()
			minNode.left = n.left
			n = minNode
		}else if n.left != nil{
			n = n.left
		}else {
			n = n.right
		}
	}
	if n != nil{
		n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
		n = n.balance()
	}
	return n


}

func (n *AvlNode) getMinNode2() *AvlNode{
	if n == nil || n.left == nil{
		return n
	}
	return n.left.getMinNode()
}

func (n *AvlNode) getMinNode() *AvlNode{
	if n == nil {
		return nil
	}
	for n.left != nil{
		n = n.left
	}
	return n
}


func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}
