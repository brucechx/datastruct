package simp_avl

import (
	"testing"
	"fmt"
	"strings"
	"github.com/stretchr/testify/assert"
)

/*
				8
			  /    \
             4      10
           /  \     / \
         2     6   9   11
        / \   / \
       1   3 5   7
 */

func TestAvlTree_Insert_LL(t *testing.T) {
	avlTree := New()
	avlTree.Insert(8, "8")
	avlTree.Insert(4, "4")
	avlTree.Insert(10, "10")
	avlTree.Insert(3, "3")
	avlTree.Insert(9, "9")
	avlTree.Insert(2, "2")
	avlTree.Insert(1, "1")
	fmt.Println(avlTree.root.key, " ", avlTree.root.getHeight())
	printAvlTree(avlTree)
	/*
		   8	 		    8
		  /  \	 左左	   /  \
		 4	 10	 ==> 	  3    10
		/	 /			 / \   /
	   3(a) 9			2   4 9
      /                /
	 2                1
	/
   1
	 */
}

func TestAvlTree_Insert_LL_2(t *testing.T) {
	avlTree := New()
	avlTree.Insert(8, "8")
	avlTree.Insert(4, "4")
	avlTree.Insert(5, "5")
	fmt.Println(avlTree.root.key, " ", avlTree.root.getHeight())
	printAvlTree(avlTree)
	/*
		   8 (a)	 	    5
		  / 	 左左	   /  \
		 4		 ==> 	  4    8
		  \
	       5

	 */
}

func TestAvlTree_Insert_RR(t *testing.T) {
	avlTree := New()
	avlTree.Insert(8, "8")
	avlTree.Insert(10, "10")
	avlTree.Insert(11, "11")
	fmt.Println(avlTree.root.key, " ", avlTree.root.getHeight())
	printAvlTree(avlTree)
	/*
		   8(a)	 		    10
		    \	 右右	   /  \
		 	10	 ==> 	  8    11
			  \
	    	   11
	 */
}

func TestAvlTree_Insert_LR(t *testing.T) {
	avlTree := New()
	avlTree.Insert(8, "8")
	avlTree.Insert(3, "3")
	avlTree.Insert(10, "10")
	avlTree.Insert(2, "2")
	avlTree.Insert(4, "4")
	avlTree.Insert(5, "5")
	fmt.Println(avlTree.root.key, " ", avlTree.root.getHeight())
	printAvlTree(avlTree)
	/*
	      8(a)	 		     4
	     /  \	   左右	   /  \
	    3 	 10   ==> 	  3    8
	   / \               /    / \
	  2   4             2    5   10
	       \
	        5
	 */
}

func TestAvlTree_Insert_RL(t *testing.T) {
	avlTree := New()
	avlTree.Insert(6, "6")
	avlTree.Insert(3, "3")
	avlTree.Insert(10, "10")
	avlTree.Insert(8, "8")
	avlTree.Insert(7, "7")
	assert.Equal(t, 6, avlTree.root.key)
	assert.Equal(t, 2, avlTree.root.getHeight())
	printAvlTree(avlTree)
	/*
	       6(a)              6
	     /   \	   右左	   /   \
	    3 	  10   ==> 	  3     8
	         /                 /  \
	        8                 7    10
	       /
	      7
	 */
}


func printAvlTree(tree *AvlTree){
	l := tree.PreOrder()  // 根左右
	var preData []string
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*Node)
		preData = append(preData, fmt.Sprintf("%v", obj.value))
	}
	fmt.Println("pre data=", strings.Join(preData, ","))
	l = tree.InOrder()  //  左根右
	var dataSlice []string
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*Node)
		dataSlice = append(dataSlice, fmt.Sprintf("%v", obj.value))
	}
	fmt.Println("mid data=", strings.Join(dataSlice, ","))

	l = tree.SeqTraverse()
	var seqSlice []string
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*Node)
		seqSlice = append(seqSlice, fmt.Sprintf("%v", obj.value))
	}
	fmt.Println("seq data=", strings.Join(seqSlice, ","))

}


