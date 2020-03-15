package binary_tree

/*
	基于前序遍历和中序遍历；重建二叉树
*/
/*
	前序遍历： 根 ｜ 左｜ 右 // 1, 2, 3
    中序遍历： 左 ｜ 根｜ 右 // 2, 1, 3 == index := 1
*/

func RebuildTree(preOrderResult, inOrderResult []int) (root *BinTreeNode){
	if len(preOrderResult) == 0{
		return nil
	}
	root = &BinTreeNode{
		data:   preOrderResult[0],
	}
	rootIndex := getIndex(inOrderResult, preOrderResult[0])
	root.lChild = RebuildTree(preOrderResult[1:rootIndex+1], inOrderResult[:rootIndex])
	root.rChild = RebuildTree(preOrderResult[rootIndex+1:], inOrderResult[rootIndex+1:])
	return root
}



func RebuildTree2(preOrderResult, inOrderResult []int) (root *BinTreeNode){
	if len(preOrderResult) == 0{
		return nil
	}
	return rebuildTree(preOrderResult, inOrderResult,
		0, len(preOrderResult) - 1,
		0, len(inOrderResult) - 1)
}



func rebuildTree(preOrderResult, inOrderResult []int,
	preOrderStartIndex, preOrderEndIndex int,
	inOrderStartIndex, inOrderEndIndex int)(rootNode *BinTreeNode){
	preOrderRootIndex := preOrderStartIndex
	rootNodeVal := preOrderResult[preOrderStartIndex]
	rootNode = &BinTreeNode{
		lChild:  nil,
		rChild: nil,
		data: rootNodeVal,
	}
	if preOrderStartIndex == preOrderEndIndex{
		return
	}
	// to get left tree pre order end index
	inOrderRootIndex := getIndex(inOrderResult, rootNodeVal)
	if inOrderRootIndex == -1{
		return
	}
	// 边界
	leftTreeNodeNums := inOrderRootIndex - inOrderStartIndex
	if leftTreeNodeNums > 0 {
		// 完善left 和 right 子树
		leftTreePreOrderStartIndex := preOrderRootIndex + 1
		leftTreePreOrderEndIndex := preOrderRootIndex + leftTreeNodeNums
		leftTreeInOrderStartIndex := inOrderStartIndex
		leftTreeInOrderEndIndex := inOrderRootIndex - 1
		//左子树
		rootNode.lChild = rebuildTree(preOrderResult, inOrderResult,
			leftTreePreOrderStartIndex, leftTreePreOrderEndIndex,
			leftTreeInOrderStartIndex, leftTreeInOrderEndIndex)
	}
	rightTreeNodeNums := inOrderEndIndex - inOrderRootIndex
	if rightTreeNodeNums > 0{
		// 右子树
		rightTreePreOrderStartIndex := preOrderEndIndex - rightTreeNodeNums + 1
		rightTreePreOrderEndIndex := preOrderEndIndex
		rightTreeInOrderStartIndex := inOrderRootIndex + 1
		rightTreeInOrderEndIndex := inOrderEndIndex
		rootNode.rChild = rebuildTree(preOrderResult, inOrderResult,
			rightTreePreOrderStartIndex, rightTreePreOrderEndIndex,
			rightTreeInOrderStartIndex, rightTreeInOrderEndIndex)
	}
	return
}

func getIndex(input []int, target int) int{
	for i, v := range input{
		if v == target{
			return i
		}
	}
	return -1
}