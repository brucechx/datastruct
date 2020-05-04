package Bellman_Ford

import (
	"testing"
	"fmt"
)

func TestBellmanFord(t *testing.T) {
	// graph
	g := NewDiGraph()
	g.Insert(1, 2, 2)
	g.Insert(1, 3, 1)
	g.Insert(2, 4, 2)
	g.Insert(3, 5, 1)
	g.Insert(5, 4, 1)

	// start node
	s := 1
	dists := BellmanFord(g, s)
	fmt.Println(dists)
	fmt.Println(dists[4])
}
