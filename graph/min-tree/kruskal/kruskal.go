package kruskal

import (
	"math"
	"fmt"
)

func (g *Graph)Kruskal(){
	var UNodes, VNodes []Node
	for _, v := range g.nodes{
		UNodes = append(UNodes, *v)
	}

	for len(VNodes) < len(UNodes){
		var pre, next Node
		distance := math.MaxInt8
		for _, s := range UNodes{
			for _, d  := range UNodes{
				if isExistInNode(VNodes, s) && isExistInNode(VNodes, d) || s == d{
					continue
				}
				w := g.edges[s][d]
				if w > 0 && w < distance{
					distance = w
					pre = s
					next = d
				}
			}
		}
		if ! isExistInNode(VNodes, pre){
			VNodes = append(VNodes, pre)
		}
		if ! isExistInNode(VNodes, next){
			VNodes = append(VNodes, next)
		}
	}
	str := ""
	for _, iNode := range VNodes {
		str += iNode.String() + " -> "
	}
	fmt.Println(str)
}
