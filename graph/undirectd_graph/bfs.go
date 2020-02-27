package undirectd_graph

// 广度遍历
// 图的遍历，深度优先与广度优先遍历
// 首先bfs 广度优先搜索
// 此处结合队列实现图的广度优先遍历

func (g *Graph) BFS(f func(n *Node)){
	allNodes := NewQueue()
	firstNode := g.nodes[0]
	allNodes.EnQueue(firstNode)

	visitMap := make(map[*Node]bool)
	visitMap[firstNode] = true

	for ! allNodes.IsEmpty(){
		qNode := allNodes.DeQueue()
		node := qNode.(*Node)
		if ! visitMap[node]{
			visitMap[node] = true
		}
		for _, eNode := range g.edges[node]{
			if visitMap[eNode]{
				continue
			}
			allNodes.EnQueue(eNode)
			visitMap[eNode] = true
		}
		if f != nil{
			f(node)
		}
	}
}
