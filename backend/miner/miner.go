package miner

import (
    bc "blockchain-visualizer/blockchain"  // Alias the package to avoid confusion
    "fmt"
    "sync"
    "time"
)

// Miner mines a new block
func Miner(blockchain *bc.Blockchain, transactions []bc.Transaction, difficulty int, wg *sync.WaitGroup, resultChan chan *bc.Block) {
    defer wg.Done()

    lastBlock := blockchain.GetLatestBlock()
    newBlock := bc.NewBlock(lastBlock.Index+1, lastBlock.Hash, transactions)  // Call the package function

    for {
        newBlock.Hash = newBlock.CalculateHash()
        if newBlock.IsValidHash() {
            fmt.Printf("Miner found a valid block with nonce: %d\n", newBlock.Nonce)
            resultChan <- newBlock
            return
        }
        newBlock.Nonce++
    }
}

// StartMining starts multiple miners concurrently
func StartMining(blockchain *bc.Blockchain, transactions []bc.Transaction, difficulty int, numMiners int) *bc.Block {
    var wg sync.WaitGroup
    resultChan := make(chan *bc.Block, numMiners)
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
        return nil // Return nil to indicate timeout
    }
}