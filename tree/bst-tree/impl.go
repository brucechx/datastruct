package bst_tree


/*
	子结点值都大于父节点
*/

func (b *BstTree) InsertImpl(key int, i Item){
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
	b.insertImpl(b.root, newNode)
}

func (b *BstTree) insertImpl(curr, new *Node) {

}