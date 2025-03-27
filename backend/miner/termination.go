package miner

import (
    "fmt"
    "sync"
)

// Color represents token color in the termination algorithm
type Color int

const (
    White Color = iota
    Black
)

// Node represents a node in the spanning tree
type Node struct {
    ID       int
    Parent   *Node
    Children []*Node
    Active   bool
    Color    Color
    mutex    sync.Mutex
}

// SpanningTree represents a spanning tree for termination detection
type SpanningTree struct {
    Root      *Node
    NodeCount int
}

// NewSpanningTree creates a new spanning tree with n nodes
func NewSpanningTree(n int) *SpanningTree {
    fmt.Println("➤ Created spanning tree with", n, "nodes")

    root := &Node{
        ID:       0,
        Active:   true,
        Color:    White,
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
            Color:    White,
            Children: make([]*Node, 0),
        }
        parent.Children = append(parent.Children, node)
        nodes[i] = node
    }

    return st
}

// MarkNodeTerminated marks a node as terminated (inactive)
func (st *SpanningTree) MarkNodeTerminated(nodeID int) {
    // Find the node with the given ID
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

    if node != nil {
        node.mutex.Lock()
        node.Active = false
        node.mutex.Unlock()
        fmt.Printf("➤ Node %d marked inactive\n", node.ID)
    }
}

// DetectTermination initiates termination detection algorithm
func (st *SpanningTree) DetectTermination() bool {
    fmt.Println("➤ Starting termination detection...")
    
    // Initialize all nodes to white
    st.initializeColors(st.Root)
    
    // Send white token down the tree
    terminated := st.sendToken(st.Root)
    
    if terminated {
        fmt.Println("➤ Termination detection completed: All processes have terminated")
    } else {
        fmt.Println("➤ Termination detection completed: Some processes still active")
    }
    
    return terminated
}

// initializeColors sets all nodes to white
func (st *SpanningTree) initializeColors(node *Node) {
    node.mutex.Lock()
    node.Color = White
    node.mutex.Unlock()
    
    for _, child := range node.Children {
        st.initializeColors(child)
    }
}

// sendToken sends a token through the tree
// Returns true if termination is detected, false otherwise
func (st *SpanningTree) sendToken(node *Node) bool {
    // First check if current node is active
    node.mutex.Lock()
    active := node.Active
    node.mutex.Unlock()
    
    if active {
        fmt.Printf("➤ Node %d is still active, termination not complete\n", node.ID)
        return false
    }
    
    fmt.Printf("➤ Sending token to node %d\n", node.ID)
    
    // Token color remains white unless a black node is found
    tokenColor := White
    
    // Send token to children
    for _, child := range node.Children {
        childTerminated := st.sendToken(child)
        
        if !childTerminated {
            return false // If any subtree is not terminated, return false immediately
        }
        
        // Check child's color
        child.mutex.Lock()
        if child.Color == Black {
            tokenColor = Black // Token becomes black if any child is black
        }
        child.mutex.Unlock()
    }
    
    // Update node's color based on token
    node.mutex.Lock()
    prevColor := node.Color
    node.Color = tokenColor
    node.mutex.Unlock()
    
    fmt.Printf("➤ Node %d processed token: was %v, now %v\n", 
        node.ID, colorToString(prevColor), colorToString(tokenColor))
    
    // If this is the root and token is white, termination is detected
    if node.Parent == nil && tokenColor == White {
        return true
    }
    
    return true
}

// Helper function to convert color to string
func colorToString(c Color) string {
    if c == White {
        return "White"
    }
    return "Black"
}

// PrintTreeStatus prints the current status of all nodes in the tree
func (st *SpanningTree) PrintTreeStatus() {
    fmt.Println("➤ Current tree status:")
    st.printNodeStatus(st.Root, 0)
}

// printNodeStatus prints the status of a node and its children
func (st *SpanningTree) printNodeStatus(node *Node, level int) {
    indent := ""
    for i := 0; i < level; i++ {
        indent += "  "
    }
    
    node.mutex.Lock()
    active := node.Active
    color := node.Color
    node.mutex.Unlock()
    
    fmt.Printf("%sNode %d - Active: %t, Color: %s\n", 
        indent, node.ID, active, colorToString(color))
    
    for _, child := range node.Children {
        st.printNodeStatus(child, level+1)
    }
}