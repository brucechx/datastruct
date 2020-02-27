package avl_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestAvlTree_Insert_Left_Rotate(t *testing.T) {
	avl := NewAvlTree()
	avl.Insert(100, 100)
	avl.Insert(60, 60)
	avl.Insert(120, 120)
	avl.Insert(110, 110)
	avl.Insert(130, 130)
	assert.Equal(t, "100,;60,120,;110,130,;", avl.String())
	avl.Insert(140, 140)
	assert.Equal(t, "120,;100,130,;60,110,140,;", avl.String())
	avl.Delete(100)
	fmt.Println(avl)
}

