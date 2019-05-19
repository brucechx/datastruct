package imm_avl

type nodes []*node

func (ns nodes) reset() {
	for i := range ns {
		ns[i] = nil
	}
}

type node struct {
	balance  int8 // bounded, |balance| should be <= 1
	children [2]*node
	entry    Entry
}

// copy returns a copy of this node with pointers to the original
// children.
func (n *node) copy() *node {
	return &node{
		balance:  n.balance,
		children: [2]*node{n.children[0], n.children[1]},
		entry:    n.entry,
	}
}

// newNode returns a new node for the provided entry.  A nil
// entry is used to represent the dummy node.
func newNode(entry Entry) *node {
	return &node{
		entry:    entry,
		children: [2]*node{},
	}
}
