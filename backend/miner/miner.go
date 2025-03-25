package miner

import (
	"fmt"
	"sync"
	"time"
	"blockchain-web-app/backend/blockchain"
)

// Miner mines a new block
func Miner(blockchain *blockchain.Blockchain, transactions []blockchain.Transaction, difficulty int, wg *sync.WaitGroup, resultChan chan blockchain.Block) {
	defer wg.Done()

	lastBlock := blockchain.Chain[len(blockchain.Chain)-1]
	newBlock := blockchain.NewBlock(lastBlock.Index+1, transactions, lastBlock.Hash)

	for {
		newBlock.Nonce++
		newBlock.Hash = newBlock.CalculateHash()
		if newBlock.IsValidHash(difficulty) {
			fmt.Printf("Miner found a valid block with nonce: %d\n", newBlock.Nonce)
			resultChan <- newBlock
			return
		}
	}
}

// StartMining starts multiple miners concurrently
func StartMining(blockchain *blockchain.Blockchain, transactions []blockchain.Transaction, difficulty int, numMiners int) blockchain.Block {
	var wg sync.WaitGroup
	resultChan := make(chan blockchain.Block, numMiners)
	timeout := time.After(10 * time.Second) // 10-second timeout

	for i := 0; i < numMiners; i++ {
		wg.Add(1)
		go Miner(blockchain, transactions, difficulty, &wg, resultChan)
	}

	// Wait for the first miner to find a valid block or timeout
	select {
	case validBlock := <-resultChan:
		wg.Wait()
		return validBlock
	case <-timeout:
		fmt.Println("Mining timed out!")
		wg.Wait()
		return blockchain.Block{} // Return an empty block to indicate timeout
	}
}