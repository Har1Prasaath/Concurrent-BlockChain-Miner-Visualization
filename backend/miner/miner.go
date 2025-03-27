package miner

import (
	bc "blockchain-visualizer/blockchain"
	"fmt"
	"sync"
	"time"
)

// Miner mines a new block with improved termination detection
func Miner(blockchain *bc.Blockchain, transactions []bc.Transaction, difficulty int,
	wg *sync.WaitGroup, resultChan chan *bc.Block, st *SpanningTree, nodeID int) {
	defer wg.Done()

	lastBlock := blockchain.GetLatestBlock()
	newBlock := bc.NewBlock(lastBlock.Index+1, lastBlock.Hash, transactions)

	// Get termination channel for quick notification
	terminationChan := st.GetTerminationChannel()

	// Mining loop with faster termination check
	for {
		// Check for termination before each iteration
		if st.IsTerminated() {
			return
		}

		// Try to mine the block
		newBlock.Hash = newBlock.CalculateHash()
		if newBlock.IsValidHash() {
			fmt.Printf("Miner %d found a valid block with nonce: %d\n", nodeID, newBlock.Nonce)

			// Signal termination immediately
			st.TerminateNode(nodeID)

			// Try to send the result, but with timeout to avoid deadlock
			select {
			case resultChan <- newBlock:
				// Successfully sent
			case <-time.After(200 * time.Millisecond):
				// Timed out - another miner already found a block
			}
			return
		}

		// Periodically check for termination again
		if newBlock.Nonce%1000 == 0 {
			select {
			case <-terminationChan:
				return // Termination signal received
			default:
				// Continue mining
			}
		}

		newBlock.Nonce++
	}
}

// Helper function to check if a node is active
func isNodeActive(st *SpanningTree, nodeID int) bool {
	st.mutex.RLock()
	defer st.mutex.RUnlock()

	// Find the node (simplified version - in real code you might want to optimize this)
	queue := []*Node{st.Root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.ID == nodeID {
			current.mutex.Lock()
			active := current.Active
			current.mutex.Unlock()
			return active
		}

		queue = append(queue, current.Children...)
	}

	return false // Node not found, consider inactive
}

// StartMining starts multiple miners concurrently with improved termination detection
func StartMining(blockchain *bc.Blockchain, transactions []bc.Transaction, difficulty int, numMiners int) *bc.Block {
	var wg sync.WaitGroup
	resultChan := make(chan *bc.Block, 1)   // Buffer size 1 to prevent blocking
	timeout := time.After(10 * time.Second) // 10-second timeout

	// Create the spanning tree for termination detection
	spanningTree := NewSpanningTree(numMiners)

	// Start all miners
	for i := 0; i < numMiners; i++ {
		wg.Add(1)
		go Miner(blockchain, transactions, difficulty, &wg, resultChan, spanningTree, i)
	}

	// Wait for the first miner to find a valid block or timeout
	var validBlock *bc.Block

	select {
	case block := <-resultChan:
		fmt.Println("Valid block found, terminating all miners...")
		validBlock = block
	case <-timeout:
		fmt.Println("Mining timed out!")
		// Signal termination on timeout
		spanningTree.TerminateNode(0)
	}

	// Ensure all miners have terminated before returning
	wg.Wait()
	fmt.Println("All miners have terminated.")

	return validBlock
}
