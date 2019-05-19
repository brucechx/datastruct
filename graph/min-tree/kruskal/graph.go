package kruskal

import (
	"sync"
	"fmt"
)

// Node 节点元素
type Item interface {

}

// 图的顶点
type Node struct {
	value Item
}

// 图
type Graph struct {
	mutex sync.RWMutex  // 保证线程安全
	nodes []*Node          // 节点集
	edges map[Node]map[Node]int // 邻接表表示的无向图
	v, e   int
}

// 图中顶点个数
func (g *Graph)V() int{
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.v
}

// 图中边的个数
func (g *Graph)E() int{
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.e
}

// 增加顶点
func (g *Graph) AddNode (n *Node){
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if isExistInSlice(g.nodes, n){
		return
	}
	g.nodes = append(g.nodes, n)
	g.v++
}

// 增加边
func (g *Graph) AddEdge(u, v *Node, w int){
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[Node]map[Node]int)
	}
	if g.edges[*u] == nil{
		g.edges[*u] = make(map[Node]int)
	}
	if g.edges[*v] == nil{
		g.edges[*v] = make(map[Node]int)
	}
	g.edges[*u][*v] = w // 建立 u->v 的边
	g.edges[*v][*u] = w // 由于是无向图，同时存在 v->u 的边
	g.e++
}

// Degree returns the number of vertices connected to u
func (g *Graph) Degree(u *Node) int{
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return len(g.edges[*u])
}

// 输出图
func (g *Graph) String() {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	str := ""
	for _, iNode := range g.nodes {
		str += iNode.String() + " -> "
		nexts := g.edges[*iNode]
		for next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

// 输出节点
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func isExistInSlice(tmp []*Node, t *Node) bool{
	for _, v := range tmp{
		if v.value == t.value{
			return true
		}
	}
	return false
}

func isExistInNode(tmp []Node, t Node) bool{
	for _, v := range tmp{
		if v.value == t.value{
			return true
		}
	}
	return false
}

func removeNode(tmp []Node, t Node) []Node{
	var res []Node
	for _, v := range tmp{
		if v == t{
			continue
		}
		res = append(res, v)
	}
	return res
}
