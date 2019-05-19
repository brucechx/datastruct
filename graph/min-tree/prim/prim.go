package prim

import (
	"math"
	"fmt"
)

/*
	Prim具体思想：以顶点为原则
	设图G顶点集合为U，首先任意选择图G中的一点作为起始点a，
	将该点加入集合V，再从集合U-V中找到另一点b使得点b到V中任意一点的权值最小，
	此时将b点也加入集合V；以此类推，现在的集合V={a，b}，再从集合U-V中找到另一点c使得点c到V中任意一点的权值最小，
	此时将c点加入集合V，直至所有顶点全部被加入V，此时就构建出了一颗MST。因为有N个顶点，所以该MST就有N-1条边，
	每一次向集合V中加入一个点，就意味着找到一条MST的边。
*/

func (g *Graph)Prim(root Node){
	var UNodes []Node
	for _, v := range g.nodes{
		UNodes = append(UNodes, *v)
	}
	if ! isExistInNode(UNodes, root){
		panic("node does not exist in graph")
	}
	UNodes = removeNode(UNodes, root)
	var VNodes []Node
	var next Node
	VNodes = append(VNodes, root)
	for len(UNodes) != 0{
		distance := math.MaxInt8
		for _, s := range VNodes{
			for d, w := range g.edges[s]{
				if isExistInNode(VNodes, *d) || s.value == d.value{
					continue
				}
				if w < distance{
					distance = w
					next = *d
				}
			}
		}
		VNodes = append(VNodes, next)
		UNodes = removeNode(UNodes, next)
	}
	str := ""
	for _, iNode := range VNodes {
		str += iNode.String() + " -> "
	}
	fmt.Println(str)
}