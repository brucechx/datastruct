package undirectd_graph

import (
	"errors"
	"fmt"
)

// 图结点
type Node struct {
	val Item
}

// 无向图
type Graph struct {
	nodes []*Node
	edges map[*Node][]*Node
	v int // 点数
	e int // 边数
}

func NewGraph() *Graph{
	return &Graph{
		nodes: make([]*Node, 0),
		edges: make(map[*Node][]*Node),
		v:     0,
		e:     0,
	}
}

func (g *Graph) AddNode(i Item) error{
	if g.nodeExist(i){
		return errors.New("node is exist")
	}
	node := &Node{i}
	g.nodes = append(g.nodes, node)
	g.v++
	return nil
}

func (g *Graph) AddEdge(i, j Item) error{
	srcNode := g.getNode(i)
	dstNode := g.getNode(j)
	if srcNode == nil || dstNode == nil{
		return errors.New("node is not exist")
	}
	// 无向图
	g.edges[srcNode] = append(g.edges[srcNode], dstNode)
	g.edges[dstNode] = append(g.edges[dstNode], srcNode)
	g.e ++
	return nil
}

// 图中顶点个数
func (g *Graph)V() int{
	return g.v
}

// 图中边的个数
func (g *Graph)E() int{
	return g.e
}

// Degree returns the number of vertices connected to u
func (g *Graph) Degree(u *Node) int{
	return len(g.edges[u])
}

// 输出图
func (g *Graph) String() {
	str := ""
	for _, node := range g.nodes {
		str += fmt.Sprintf("%v", node.val) + " -> "
		for _, next := range g.edges[node] {
			str += fmt.Sprintf("%v", next.val) + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

func (g *Graph) nodeExist(i Item) bool{
	v := g.getNode(i)
	if v == nil{
		return false
	}
	return true
}

func (g *Graph) getNode(i Item) *Node{
	for _, v := range g.nodes{
		if i == v.val{
			return v
		}
	}
	return nil
}


