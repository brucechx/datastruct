package binary_tree

import (
	"math"
	"fmt"
)

//二叉树节点
type BinTreeNode struct {
	data   interface{}  //数据域
	parent *BinTreeNode //父节点
	lChild *BinTreeNode //左孩子
	rChild *BinTreeNode //右孩子
	height int          //以该结点为根的子树的高度
	size   int          //该结点子孙数(包括结点本身)
}

func NewBinTreeNode(e interface{}) *BinTreeNode {
	return &BinTreeNode{data: e, size: 1}
}

//获得节点数据
func (b *BinTreeNode) GetData() interface{} {
	if b == nil {
		return nil
	}
	return b.data
}

//设置节点数据
func (b *BinTreeNode) SetData(e interface{}) {
	b.data = e
}

//判断是否有父亲
func (b *BinTreeNode) HasParent() bool {
	return b.parent != nil
}

//获得父亲节点
func (b *BinTreeNode) GetParent() *BinTreeNode {
	if !b.HasParent() {
		return nil
	}
	return b.parent
}

//设置父亲节点
func (b *BinTreeNode) setParent(p *BinTreeNode) {
	b.parent = p
	// b.parent.SetHeight() //更新父结点及其祖先高度
	// b.parent.SetSize()   //更新父结点及其祖先规模
}

//断开与父亲的关系
func (b *BinTreeNode) CutOffParent() {
	if !b.HasParent() {
		return
	}
	if b.IsLChild() {
		b.parent.lChild = nil //断开该节点与父节点的连接
	} else {
		b.parent.rChild = nil //断开该节点与父节点的连接
	}

	b.parent = nil       //断开该节点与父节点的连接
	b.parent.SetHeight() //更新父结点及其祖先高度
	b.parent.SetSize()   //更新父结点及其祖先规模
}

//判断是否有左孩子
func (b *BinTreeNode) HasLChild() bool {
	return b.lChild != nil
}

//获得左孩子节点
func (b *BinTreeNode) GetLChild() *BinTreeNode {
	if !b.HasLChild() {
		return nil
	}
	return b.lChild
}

//设置当前结点的左孩子,返回原左孩子
func (b *BinTreeNode) SetLChild(lc *BinTreeNode) *BinTreeNode {
	oldLC := b.lChild
	if b.HasLChild() {
		b.lChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if lc != nil {
		lc.CutOffParent() //断开lc与其父结点的关系
		b.lChild = lc     //确定父子关系
		lc.setParent(b)
		b.SetHeight() //更新当前结点及其祖先高度
		b.SetSize()   //更新当前结点及其祖先规模
	}
	return oldLC
}

//判断是否有右孩子
func (b *BinTreeNode) HasRChild() bool {
	return b.rChild != nil
}

//获得右孩子节点
func (b *BinTreeNode) GetRChild() *BinTreeNode {
	if !b.HasRChild() {
		return nil
	}
	return b.rChild
}

//设置当前结点的右孩子,返回原右孩子
func (b *BinTreeNode) SetRChild(rc *BinTreeNode) *BinTreeNode {
	oldRC := b.rChild
	if b.HasRChild() {
		b.rChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if rc != nil {
		rc.CutOffParent() //断开rc与其父结点的关系
		b.rChild = rc     //确定父子关系
		rc.setParent(b)
		b.SetHeight() //更新当前结点及其祖先高度
		b.SetSize()   //更新当前结点及其祖先规模
	}
	return oldRC
}

//判断是否为叶子结点
func (b *BinTreeNode) IsLeaf() bool {
	return !b.HasLChild() && !b.HasRChild()
}

//判断是否为某结点的左孩子
func (b *BinTreeNode) IsLChild() bool {
	return b.HasParent() && b == b.parent.lChild
}

//判断是否为某结点的右孩子
func (b *BinTreeNode) IsRChild() bool {
	return b.HasParent() && b == b.parent.rChild
}

//取结点的高度,即以该结点为根的树的高度
func (b *BinTreeNode) GetHeight() int {
	return b.height
}

//更新当前结点及其祖先的高度
func (b *BinTreeNode) SetHeight() {
	newH := 0 //新高度初始化为0,高度等于左右子树高度加1中的大者
	if b.HasLChild() {
		newH = int(math.Max(float64(newH), float64(1+b.GetLChild().GetHeight())))
	}
	if b.HasRChild() {
		newH = int(math.Max(float64(newH), float64(1+b.GetRChild().GetHeight())))
	}
	if newH == b.height {
		//高度没有发生变化则直接返回
		return
	}

	b.height = newH //否则更新高度
	if b.HasParent() {
		b.GetParent().SetHeight() //递归更新祖先的高度
	}
}

//取以该结点为根的树的结点数
func (b *BinTreeNode) GetSize() int {
	return b.size
}

//更新当前结点及其祖先的子孙数
func (b *BinTreeNode) SetSize() {
	b.size = 1 //初始化为1,结点本身
	if b.HasLChild() {
		b.size += b.GetLChild().GetSize() //加上左子树规模
	}
	if b.HasRChild() {
		b.size += b.GetRChild().GetSize() //加上右子树规模
	}

	if b.HasParent() {
		b.parent.SetSize() //递归更新祖先的规模
	}

}

func (b *BinTreeNode)String() string{
	return fmt.Sprintf("node data=%v", b.data)
}