package int_heap

import (
	"fmt"
	"strings"
)

type IntHeap struct {
	Data []int
}

func NewIntHeap(data []int) *IntHeap{
	return &IntHeap{Data:data}
}

func (h *IntHeap) String() string{
	var res []string
	for _, v := range h.Data {
		res = append(res, fmt.Sprintf("%d", v))
	}
	return strings.Join(res, " ")
}

func (h *IntHeap) Size() int{
	return len(h.Data)
}

//构建堆
func (h *IntHeap)MakeSmallHeap(){
	n := h.Size()
	// 从最后一个非叶子节点开始，依次下沉父节点调整
	for i := n/2-1; i >= 0; i--{
		h.down(i, n)
	}
}


//由父节点至子节点依次建堆
//parent      : i
//left child  : 2*i+1
//right child : 2*i+2
func (h *IntHeap)down(i,n int) {
	//构建最小堆,父小于两子节点值
	for {

		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}

		//找出两个节点中最小的(less: a<b)
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && !h.Less(j1, j2) {
			j = j2 // = 2*i + 2  // right child
		}

		//然后与父节点i比较,如果父节点小于这个子节点最小值,则break,否则swap
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}


//由子节点至父节点依次重建堆
func (h *IntHeap)up(j int)  {
	for {
		i := (j - 1) / 2 // parent

		if i == j || !h.Less(j, i) {
			//less(子,父) !Less(9,5) == true
			//父节点小于子节点,符合最小堆条件,break
			break
		}
		//子节点比父节点小,互换
		h.Swap(i, j)
		j = i
	}
}

func (h *IntHeap)Push(x int){
	h.Data = append(h.Data, x)
	h.up(h.Size()-1)
	return
}

func (h *IntHeap)Pop() interface{} {
	n := h.Size() - 1
	h.Swap(0, n)
	h.down(0, n)

	old := h.Data
	n = len(old)
	x := old[n-1]
	h.Data = old[0 : n-1]
	return x
}

func (h *IntHeap)Remove(i int) interface{} {
	n := h.Size() - 1
	if n != i {
		h.Swap(i, n)
		h.down(i, n)
		h.up(i)
	}
	return h.Pop()
}

func (h *IntHeap)Less(a, b int)bool{
	//升序 Less(heap[a] > heap[b])	//最大堆
	//降序 Less(heap[a] < heap[b])	//最小堆
	return h.Data[a] < h.Data[b]
}

func (h *IntHeap)Swap(a,b int){
	h.Data[a], h.Data[b] = h.Data[b], h.Data[a]
}

func (h *IntHeap)HeapSort(){
	for i := h.Size()-1 ;i > 0;i--{
		//移除顶部元素到数组末尾,然后剩下的重建堆,依次循环
		h.Swap(0, i)
		h.down(0, i)
	}
}
