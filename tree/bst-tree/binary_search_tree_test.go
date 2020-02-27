package bst_tree

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
// 拓扑图1
			6
	     /     \
	  1         7
		\        \
         5        9
        /        / \
       3        8   10
      / \
     2   4
*/

/*
// 拓扑图2
			6
	     /     \
	  1         7
		\        \
         5        9
        /        / \
       3        8   11
      / \			/
     2   4         10
*/

func initBinarySearchTree() *BstTree {
	bst := NewBstTree(6, 6)
	bst.Insert(1, 1)
	bst.Insert(7, 7)
	bst.Insert(5, 5)
	bst.Insert(9, 9)
	bst.Insert(3, 3)
	bst.Insert(8, 8)
	bst.Insert(10, 10)
	bst.Insert(2, 2)
	bst.Insert(4, 4)
	return bst
}

func initBinarySearchTree2() *BstTree {
	bst := NewBstTree(6, 6)
	bst.Insert(1, 1)
	bst.Insert(7, 7)
	bst.Insert(5, 5)
	bst.Insert(9, 9)
	bst.Insert(3, 3)
	bst.Insert(8, 8)
	bst.Insert(11, 11)
	bst.Insert(2, 2)
	bst.Insert(4, 4)
	bst.Insert(10, 10)
	return bst
}

func TestBstTree(t *testing.T){
	TestBstTree_PreOrder(t)
	TestBstTree_MidOrder(t)
	TestBstTree_PostOrder(t)
	TestBstTree_Search(t)
	TestBstTree_Exist(t)
	TestBstTree_MaxNode_MinNode(t)
	TestBstTree_PreSuccessor(t)
	TestBstTree_Successor(t)
	TestBstTree_Remove1_1(t)
	TestBstTree_Remove1_2(t)
	TestBstTree_Remove1_3(t)
	TestBstTree_Remove1_4(t)
	TestBstTree_Remove2(t)
}

func TestBstTree_PreOrder(t *testing.T) {
	bstTree := initBinarySearchTree()
	var result []string
	bstTree.PreOrder(func(i Item) {
		result = append(result, fmt.Sprintf("%v", i))
	})
	res := strings.Join(result, ",")
	assert.Equal(t, "6,1,5,3,2,4,7,9,8,10", res)
}

func TestBstTree_MidOrder(t *testing.T) {
	bstTree := initBinarySearchTree()
	var result []string
	bstTree.MidOrder(func(i Item) {
		result = append(result, fmt.Sprintf("%v", i))
	})
	res := strings.Join(result, ",")
	assert.Equal(t, "1,2,3,4,5,6,7,8,9,10", res)
}

func TestBstTree_PostOrder(t *testing.T) {
	bstTree := initBinarySearchTree()
	var result []string
	bstTree.PostOrder(func(i Item) {
		result = append(result, fmt.Sprintf("%v", i))
	})
	res := strings.Join(result, ",")
	assert.Equal(t, "2,4,3,5,1,8,10,9,7,6", res)
}

func TestBstTree_Search(t *testing.T) {
	bstTree := initBinarySearchTree()
	val, err := bstTree.Search(3)
	assert.Equal(t, nil, err)
	assert.Equal(t, val, 3)
}

func TestBstTree_Exist(t *testing.T) {
	bstTree := initBinarySearchTree()
	res := bstTree.Exist(3)
	assert.Equal(t, true, res)
	res = bstTree.Exist(33)
	assert.Equal(t, false, res)
}

func TestBstTree_MaxNode_MinNode(t *testing.T) {
	bstTree := initBinarySearchTree()
	maxNode := bstTree.MaxNode()
	assert.Equal(t, 10, maxNode)
	minNode := bstTree.MinNode()
	assert.Equal(t, 1, minNode)
}

func TestBstTree_PreSuccessor(t *testing.T) {
	bstTree := initBinarySearchTree()
	v1, _ := bstTree.PreSuccessor(4)
	assert.Equal(t, 3, v1)
	v2, _ := bstTree.PreSuccessor(2)
	assert.Equal(t, 1, v2)
	v3, _ := bstTree.PreSuccessor(6)
	assert.Equal(t, 5, v3)
}

func TestBstTree_Successor(t *testing.T) {
	bstTree := initBinarySearchTree()
	v1, _ := bstTree.Successor(7)
	assert.Equal(t, 8, v1)
	v2, _ := bstTree.Successor(5)
	assert.Equal(t, 6, v2)
	v3, _ := bstTree.Successor(2)
	assert.Equal(t, 3, v3)
}

func TestBstTree_Remove1_1(t *testing.T) {
	bstTree := initBinarySearchTree()
	err := bstTree.Remove(2) // 3.1
	assert.Equal(t, nil, err)
	res := bstTree.String()
	assert.Equal(t, "1,3,4,5,6,7,8,9,10", res)

}

func TestBstTree_Remove1_2(t *testing.T) {
	bstTree := initBinarySearchTree()
	err := bstTree.Remove(1) // 3.2
	assert.Equal(t, nil, err)
	res := bstTree.String()
	assert.Equal(t, "2,3,4,5,6,7,8,9,10", res)
}

func TestBstTree_Remove1_3(t *testing.T) {
	bstTree := initBinarySearchTree()
	err := bstTree.Remove(9) // 3.3.1
	assert.Equal(t, nil, err)
	res := bstTree.String()
	assert.Equal(t, "1,2,3,4,5,6,7,8,10", res)
}

func TestBstTree_Remove1_4(t *testing.T) {
	bstTree2 := initBinarySearchTree2()
	err := bstTree2.Remove(9) // 3.3.2
	assert.Equal(t, nil, err)
	res := bstTree2.String()
	assert.Equal(t, "1,2,3,4,5,6,7", res)
}

func TestBstTree_Remove2(t *testing.T) {
	bstTree := initBinarySearchTree()
	bstTree.Remove2(2)
	res := bstTree.String()
	assert.Equal(t, "1,3,4,5,6,7,8,9,10", res)
}
