package miner

import (
    bc "blockchain-visualizer/blockchain"
    "fmt"
    "sync"
    "time"
)

// Miner mines a new block
func Miner(blockchain *bc.Blockchain, transactions []bc.Transaction, difficulty int,
    wg *sync.WaitGroup, resultChan chan *bc.Block, stopChan <-chan struct{}, minerID int, doneChan chan<- int) {
    defer wg.Done()

    lastBlock := blockchain.GetLatestBlock()
    newBlock := bc.NewBlock(lastBlock.Index+1, lastBlock.Hash, transactions)

    // Mining loop
    for {
        // Check for stop signal
        select {
        case <-stopChan:
            fmt.Printf("◆ Miner %d stopping\n", minerID)
            doneChan <- minerID // Signal that this miner has stopped
            return
        default:
            // Continue mining
        }

        // Try to mine the block
        newBlock.Hash = newBlock.CalculateHash()
        if newBlock.IsValidHash() {
            fmt.Printf("◆ Miner %d found valid block with nonce: %d\n", minerID, newBlock.Nonce)
            resultChan <- newBlock
            doneChan <- minerID // Signal that this miner has stopped
            return
        }

        newBlock.Nonce++
    }
}

// StartMining starts multiple miners concurrently
func StartMining(blockchain *bc.Blockchain, transactions []bc.Transaction, difficulty int, numMiners int) *bc.Block {
    fmt.Printf("▶ Started mining with %d concurrent miners\n", numMiners)
    
    var wg sync.WaitGroup
    resultChan := make(chan *bc.Block, 1)
    stopChan := make(chan struct{})
    doneChan := make(chan int, numMiners) // Channel to track terminated miners
    
    spanningTree := NewSpanningTree(numMiners)
    timeout := time.After(10 * time.Second)

    // Start all miners
    for i := 0; i < numMiners; i++ {
        wg.Add(1)
        go Miner(blockchain, transactions, difficulty, &wg, resultChan, stopChan, i, doneChan)
    }

    // Wait for result or timeout
    var validBlock *bc.Block

    select {
    case block := <-resultChan:
        validBlock = block
        close(stopChan) // Signal all miners to stop
        
    case <-timeout:
        fmt.Println("▶ Mining timed out after 10 seconds")
        close(stopChan) // Signal all miners to stop on timeout
    }

    // Wait for all miners to terminate and mark them in the spanning tree
    terminatedCount := 0
    timeoutLoop := false
    for terminatedCount < numMiners && !timeoutLoop {
        select {
        case minerID := <-doneChan:
            spanningTree.MarkNodeTerminated(minerID)
            terminatedCount++
        case <-time.After(5 * time.Second):
            fmt.Println("▶ Timed out waiting for miners to terminate")
            timeoutLoop = true
        }
    }

    // Run termination detection algorithm
    allTerminated := spanningTree.DetectTermination()
    if allTerminated {
        fmt.Println("▶ All miners have successfully terminated")
    } else {
        fmt.Println("▶ Some miners did not terminate properly")
    }

    wg.Wait()
    return validBlock
}