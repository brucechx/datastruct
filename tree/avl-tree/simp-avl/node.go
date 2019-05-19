package simp_avl

// Node a single node that composes the tree
type Node struct {
	key   int  // 中序遍历的节点序号
	value interface{} // 节点存储的值
	left  *Node //left
	right *Node //right
	height int // 高度
}

func (node *Node) getHeight() int {
	if node != nil {
		return node.height
	}
	return -1
}


