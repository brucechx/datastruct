package bfs_dfs

// 深度遍历
func (g *Graph) DFS(f func(n *Node)){
	allNodes := NewStack()
	firstNode := g.nodes[0]
	allNodes.Push(firstNode)

	visitMap := make(map[*Node]bool)
	visitMap[firstNode] = true

	for ! allNodes.IsEmpty(){
		sNode := allNodes.Pop()
		node := sNode.(*Node)
		if ! visitMap[node] {
			visitMap[node] = true
		}
		for _, eNode := range g.edges[node]{
			if visitMap[eNode]{
				continue
			}
			allNodes.Push(eNode)
			visitMap[eNode] = true
		}
		if f != nil{
			f(node)
		}
	}
}
