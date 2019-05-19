package undirectd_graph

import "sync"

// 图的遍历，深度优先与广度优先遍历
// 首先bfs 广度优先搜索
// 此处结合队列实现图的广度优先遍历

func (g *Graph)BFS(f func(node *Node)){
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// 初始化队列
	q := NodeQueue{}
	q.New()
	// 取出图中的第一个节点，放入队列
	n := g.nodes[0]
	q.EnQueue(*n)
	// 初始化节点列表，以表示访问过
	visited := make(map[*Node]bool)
	visited[n] = true

	for {
		if q.IsEmpty(){
			break
		}
		node := q.DeQueue()
		if !visited[node]{
			visited[node] = true
		}
		for _, v := range g.edges[*node]{
			if visited[v]{
				continue
			}
			q.EnQueue(*v)
			visited[v] = true
		}
		if f != nil{
			f(node)
		}
	}

}

type NodeQueue struct {
	nodes []Node
	lock  sync.RWMutex
}

func(q *NodeQueue) New() *NodeQueue{
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = []Node{}
	return q
}

// 入队列
func (q *NodeQueue) EnQueue(n Node){
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = append(q.nodes, n)
}

// 出队列
func (q *NodeQueue) DeQueue() *Node{
	q.lock.Lock()
	defer q.lock.Unlock()

	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return &node
}

// 判断队列是否为空
func (q *NodeQueue) IsEmpty() bool{
	q.lock.RLock()
	defer q.lock.RUnlock()
	return 0 == len(q.nodes)
}