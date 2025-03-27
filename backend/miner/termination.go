package miner

import (
    "sync"
    "sync/atomic"
)

// GlobalTermination provides a fast way to signal termination to all miners
type GlobalTermination struct {
    terminated int32 // Using atomic operations for thread safety
}

// NewGlobalTermination creates a new termination controller
func NewGlobalTermination() *GlobalTermination {
    return &GlobalTermination{
        terminated: 0,
    }
}

// Terminate signals termination to all miners
func (gt *GlobalTermination) Terminate() {
    atomic.StoreInt32(&gt.terminated, 1)
}

// IsTerminated checks if mining should be terminated
func (gt *GlobalTermination) IsTerminated() bool {
    return atomic.LoadInt32(&gt.terminated) == 1
}

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
    Root         *Node
    NodeCount    int
    mutex        sync.RWMutex
    termination  *GlobalTermination
    terminationChan chan struct{}
}

// NewSpanningTree creates a new spanning tree with n nodes
func NewSpanningTree(n int) *SpanningTree {
    root := &Node{
        ID:       0,
        Active:   true,
        Children: make([]*Node, 0),
    }

    st := &SpanningTree{
        Root:         root,
        NodeCount:    n,
        termination:  NewGlobalTermination(),
        terminationChan: make(chan struct{}),
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

// TerminateNode marks a node as inactive and signals global termination
func (st *SpanningTree) TerminateNode(nodeID int) {
    st.mutex.Lock()
    defer st.mutex.Unlock()

    // Signal global termination immediately
    st.termination.Terminate()
    
    // Also close the termination channel to wake up any waiting goroutines
    select {
    case <-st.terminationChan: // Already closed
    default:
        close(st.terminationChan)
    }

    // Find the node with the given ID
    queue := []*Node{st.Root}
    var node *Node

    for len(queue) > 0 && node == nil {
        current := queue[0]
        queue = queue[1:]

        if current.ID == nodeID {
            node = current
            break
        }

        queue = append(queue, current.Children...)
    }

    if node != nil {
        // Mark the node as inactive
        node.mutex.Lock()
        node.Active = false
        node.mutex.Unlock()
        
        // Propagate termination upward
        st.propagateTermination(node)
    }
}

// propagateTermination propagates termination up the tree
func (st *SpanningTree) propagateTermination(node *Node) {
    // If this is the root, we're done
    if node.Parent == nil {
        return
    }

    // Check if all siblings are inactive
    allSiblingsInactive := true
    
    for _, child := range node.Parent.Children {
        child.mutex.Lock()
        active := child.Active
        child.mutex.Unlock()
        
        if active {
            allSiblingsInactive = false
            break
        }
    }
    
    // If all siblings are inactive, mark parent as inactive and continue up
    if allSiblingsInactive {
        node.Parent.mutex.Lock()
        node.Parent.Active = false
        node.Parent.mutex.Unlock()
        
        st.propagateTermination(node.Parent)
    }
}

// IsTerminated provides quick access to check if mining should terminate
func (st *SpanningTree) IsTerminated() bool {
    return st.termination.IsTerminated()
}

// GetTerminationChannel returns a channel that is closed when termination occurs
func (st *SpanningTree) GetTerminationChannel() <-chan struct{} {
    return st.terminationChan
}