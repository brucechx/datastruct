package prim

import "testing"

/*
          a
        /    \
     1 /  5   \2
    b   ----   e
    |   \ 4    |
   3|    \     |7
	|     \    |
    |      \   |
    c	----   d
		 6

==>  a   b   e   c   d
*/

var g Graph

func initGraph(){
	n1, n2, n3, n4, n5 := Node{"a"}, Node{"b"}, Node{"c"}, Node{"d"}, Node{"e"}

	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1, &n2, 1)
	g.AddEdge(&n1, &n5, 2)
	g.AddEdge(&n2, &n3, 3)
	g.AddEdge(&n2, &n4, 4)
	g.AddEdge(&n2, &n5, 5)
	g.AddEdge(&n3, &n4, 6)
	g.AddEdge(&n4, &n5, 7)
}

func TestPrim(t *testing.T) {
	initGraph()
	a := Node{"a"}
	g.Prim(a)

	// output:
	// a
	// b
	// e
	// c
	// d
}
