package int_heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	data := []int{100,16,4,8,70,2,36,22,5,12}
	heap := NewIntHeap(data)
	heap.MakeSmallHeap() // 构建最小堆后：
	assert.Equal(t, "2 5 4 8 12 100 36 22 16 70", heap.String())
	heap.Push(90) // 增加 90,30,1 :
	heap.Push(30)
	heap.Push(1)
	assert.Equal(t, "1 5 2 8 12 4 36 22 16 70 90 100 30", heap.String())
	n := heap.Pop() // Pop出最小值( 1 )后:
	assert.Equal(t, 1, n)
	assert.Equal(t, "2 5 4 8 12 30 36 22 16 70 90 100", heap.String())
	heap.Remove(3) // Remove()掉idx为3即值 4 后:
	assert.Equal(t, "4 5 8 16 12 30 36 22 100 70 90", heap.String())
	heap.HeapSort() // HeapSort()后:
	assert.Equal(t, "100 90 70 36 30 22 16 12 8 5 4", heap.String())
}