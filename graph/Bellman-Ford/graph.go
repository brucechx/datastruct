package Bellman_Ford

type DiGraph struct {
	outAdj    map[int]map[int]int
	inAdj     map[int]map[int]int
	nodeCount int
}

func NewDiGraph() DiGraph {
	return DiGraph{
		inAdj:     make(map[int]map[int]int),
		outAdj:    make(map[int]map[int]int),
		nodeCount: 0,
	}
}

func (g DiGraph) NodeCount() int {
	return g.nodeCount
}
func (g DiGraph) GetInbound(v int) (map[int]int) {
	Edges, ok := g.inAdj[v]
	if !ok {
		return map[int]int{}
	}
	return Edges
}

func (g *DiGraph) Insert(u, v, w int) {
	if _, ok := g.outAdj[u]; !ok {
		g.outAdj[u] = make(map[int]int)
	}
	g.outAdj[u][v] = w

	if _, ok := g.inAdj[v]; !ok {
		g.inAdj[v] = make(map[int]int)
	}
	g.inAdj[v][u] = w

	g.nodeCount++
}
