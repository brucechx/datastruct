package binary_tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRebuildBinaryTree(t *testing.T) {
	t.Run("simple three nodes tree", func(t *testing.T) {
		preOrderTraversalResult := []int{0, 1, 2}
		inOrderTraversalResult := []int{1, 0, 2}
		got := RebuildTree(preOrderTraversalResult, inOrderTraversalResult)
		left := BinTreeNode{
			lChild:nil,
			rChild:nil,
			data:1,
		}
		right := BinTreeNode{
			lChild:nil,
			rChild:nil,
			data:2,
		}
		root := BinTreeNode{
			lChild:&left,
			rChild:&right,
			data:0,
		}
		want := &root
		if !TreeEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("four nodes tree", func(t *testing.T) {
		preOrderTraversalResult := []int{0, 1, 2, 3}
		inOrderTraversalResult := []int{2, 1, 0, 3}
		got := RebuildTree(preOrderTraversalResult, inOrderTraversalResult)
		leftLeft := BinTreeNode{
			lChild:nil,
			rChild:nil,
			data:2,
		}
		left := BinTreeNode{
			lChild:&leftLeft,
			rChild:nil,
			data:1,
		}
		right := BinTreeNode{
			lChild:nil,
			rChild:nil,
			data:3,
		}
		root := BinTreeNode{
			lChild:&left,
			rChild:&right,
			data:0,
		}
		want := &root
		if !TreeEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("no node", func(t *testing.T) {
		var preOrderTraversalResult []int
		var inOrderTraversalResult []int
		got := RebuildTree(preOrderTraversalResult, inOrderTraversalResult)
		if got != nil {
			t.Errorf("got %v want %v", got, nil)
		}

	})
}

func TreeEqual(got *BinTreeNode, want *BinTreeNode) (equality bool) {
	gotTree := &binaryTree{root:got}
	wantTree := &binaryTree{root:want}
	gotResult := listString(gotTree.PreOrder())
	wantResult := listString(wantTree.PreOrder())
	if reflect.DeepEqual(gotResult, wantResult) {
		equality = true
	} else {
		fmt.Printf("got %v want %v", gotResult, wantResult)
	}

	return
}

