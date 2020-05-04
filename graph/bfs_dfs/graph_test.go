package bfs_dfs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
        1
     /     \
	2   --  5
    |   \   |
	3	--	4
*/


func initGraph() *Graph{
	g := NewGraph()

	_ = g.AddNode(1)
	_ = g.AddNode(2)
	_ = g.AddNode(3)
	_ = g.AddNode(4)
	_ = g.AddNode(5)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 5)
	_ = g.AddEdge(2, 3)
	_ = g.AddEdge(2, 4)
	_ = g.AddEdge(2, 5)
	_ = g.AddEdge(3, 4)
	_ = g.AddEdge(4, 5)
	return g
}

func TestGraph_String(t *testing.T) {
	g := initGraph()
	g.String()
}

func TestGraph_BFS(t *testing.T) {
	g := initGraph()
	var result []string
	g.BFS(func(n *Node) {
		result = append(result, fmt.Sprintf("%v", n.val))
	})
	res := strings.Join(result, ",")
	assert.Equal(t, "1,2,5,3,4", res)
}

func TestGraph_DFS(t *testing.T) {
	g := initGraph()
	var result []string
	g.DFS(func(n *Node) {
		result = append(result, fmt.Sprintf("%v", n.val))
	})
	res := strings.Join(result, ",")
	assert.Equal(t, "1,5,4,3,2", res)
}