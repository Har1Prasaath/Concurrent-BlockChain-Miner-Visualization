package network

import (
    "blockchain-visualizer/blockchain"
    "fmt"
    "time"
)

// Network simulates a peer-to-peer network
type Network struct {
    Nodes []string
}

// NewNetwork creates a new network
func NewNetwork() Network {
    return Network{
        Nodes: []string{"node1", "node2", "node3"},
    }
}

// BroadcastBlock broadcasts a mined block to all nodes
func (n Network) BroadcastBlock(block *blockchain.Block) {
    for _, node := range n.Nodes {
        fmt.Printf("Broadcasting block to %s...\n", node)
        time.Sleep(500 * time.Millisecond) // Simulate network delay
    }
    fmt.Println("Block broadcast complete!")
}