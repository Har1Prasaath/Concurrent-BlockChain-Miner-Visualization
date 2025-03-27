package miner

import (
	"sync"
)

// Node represents a node in the spanning tree
type Node struct {
	ID       int
	Parent   *Node
	Children []*Node
	Active   bool
	mutex    sync.Mutex
}

// SpanningTree represents a spanning tree for termination detection
type SpanningTree struct {
	Root      *Node
	NodeCount int
	mutex     sync.RWMutex
}

// NewSpanningTree creates a new spanning tree with n nodes
func NewSpanningTree(n int) *SpanningTree {
	root := &Node{
		ID:       0,
		Active:   true,
		Children: make([]*Node, 0),
	}

	st := &SpanningTree{
		Root:      root,
		NodeCount: n,
	}

	// Create the remaining nodes and build the tree
	nodes := make([]*Node, n)
	nodes[0] = root

	for i := 1; i < n; i++ {
		parent := nodes[(i-1)/2] // Binary tree structure for simplicity
		node := &Node{
			ID:       i,
			Parent:   parent,
			Active:   true,
			Children: make([]*Node, 0),
		}
		parent.Children = append(parent.Children, node)
		nodes[i] = node
	}

	return st
}

// TerminateNode marks a node as inactive
func (st *SpanningTree) TerminateNode(nodeID int) bool {
	st.mutex.Lock()
	defer st.mutex.Unlock()

	// Find the node with the given ID using BFS
	queue := []*Node{st.Root}
	var node *Node

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.ID == nodeID {
			node = current
			break
		}

		queue = append(queue, current.Children...)
	}

	if node == nil {
		return false
	}

	// Mark the node as inactive
	node.mutex.Lock()
	node.Active = false
	node.mutex.Unlock()

	// Check if termination should be propagated
	return st.checkTermination(node)
}

// checkTermination checks if termination should be propagated up the tree
func (st *SpanningTree) checkTermination(node *Node) bool {
	// If node has active children, can't terminate
	for _, child := range node.Children {
		child.mutex.Lock()
		active := child.Active
		child.mutex.Unlock()

		if active {
			return false
		}
	}

	// If this is the root and all children are inactive, entire computation is done
	if node.Parent == nil {
		return true
	}

	// Propagate termination to parent
	return st.TerminateNode(node.Parent.ID)
}
