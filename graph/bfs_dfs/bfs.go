package bfs_dfs

// 图的广度遍历

// 递归实现
func (g *Graph) BFS(f func(n *Node)){
	visitMap := make(map[*Node]bool)
	firstNode := g.nodes[0]
	visitMap[firstNode] = true
	f(firstNode)
	g.bfs(f, firstNode, visitMap)
}

func (g *Graph) bfs(f func(n *Node), node *Node, visitMap map[*Node]bool){
	allNodes := make([]*Node, 0)
	for _, edgeNode := range g.edges[node]{
		if _, ok := visitMap[edgeNode]; ok {
			continue
		}
		allNodes = append(allNodes, edgeNode)
		visitMap[edgeNode] = true
	}
	for _, n := range allNodes{
		if f != nil{
			f(n)
		}
	}
	for _, n := range allNodes{
		g.bfs(f, n, visitMap)
	}
}

// 非递归实现
func (g *Graph) BFS2(f func(n *Node)){
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

