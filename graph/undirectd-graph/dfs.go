package undirectd_graph

import "sync"

// 用栈实现图的深度遍历

func (g *Graph) DFS(f func(node *Node)){
	g.mutex.Lock()
	defer g.mutex.Unlock()

	s := NodeStack{}
	s.New()

	// 取图中第一个点入栈
	n := g.nodes[0]
	s.Push(*n)

	visited := make(map[*Node]bool)
	visited[n] = true

	for {
		if s.IsEmpty(){
			break
		}
		node := s.Pop()
		if !visited[node]{
			visited[node] = true
		}
		for _, v := range g.edges[*node]{
			if !visited[v]{
				s.Push(*v)
				visited[v] = true
			}
		}
		if f != nil{
			f(node)
		}
	}
}

type NodeStack struct {
	nodes []Node
	lock sync.RWMutex
}

func (s *NodeStack) New() *NodeStack{
	s.lock.Lock()
	defer s.lock.Unlock()
	s.nodes = []Node{}
	return s
}

func (s *NodeStack) Push(n Node){
	s.lock.Lock()
	defer s.lock.Unlock()
	s.nodes = append(s.nodes, n)
}

func (s *NodeStack) Pop() *Node{
	s.lock.Lock()
	defer s.lock.Unlock()

	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[0:len(s.nodes)-1]
	return &node
}

func (s *NodeStack) IsEmpty() bool{
	s.lock.Lock()
	defer s.lock.Unlock()
	return 0 == len(s.nodes)
}

