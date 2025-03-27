package blockchain

import (
	"sync"
)

type Blockchain struct {
	Blocks              []*Block
	PendingTransactions []Transaction
	mutex               sync.RWMutex // Add mutex for thread safety
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(0, "", []Transaction{})
	bc := &Blockchain{
		Blocks:              []*Block{genesisBlock},
		PendingTransactions: []Transaction{},
	}
	return bc
}

func (bc *Blockchain) AddTransaction(tx Transaction) {
	bc.mutex.Lock()         // Lock before modifying transactions
	defer bc.mutex.Unlock() // Ensure unlock happens

	bc.PendingTransactions = append(bc.PendingTransactions, tx)
}

func (bc *Blockchain) AddBlock(transactions []Transaction) *Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, prevBlock.Hash, transactions)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

func (bc *Blockchain) MinePendingTransactions(minerReward string) *Block {
	bc.mutex.Lock()

	// If we don't have any pending transactions, add just the reward transaction
	var pendingTransactionsCopy []Transaction
	if len(bc.PendingTransactions) == 0 {
		// Create empty slice but don't clear anything since there's nothing to clear
		pendingTransactionsCopy = []Transaction{}
	} else {
		// Copy pending transactions
		pendingTransactionsCopy = make([]Transaction, len(bc.PendingTransactions))
		copy(pendingTransactionsCopy, bc.PendingTransactions)
		// Clear pending transactions only after copying them
		bc.PendingTransactions = []Transaction{}
	}
	bc.mutex.Unlock()

	// Create the reward transaction
	rewardTx := NewTransaction("system", minerReward, 1.0)
	allTransactions := append(pendingTransactionsCopy, rewardTx)

	// Add the new block with all transactions
	return bc.AddBlock(allTransactions)
}

func (bc *Blockchain) GetLatestBlock() *Block {
	bc.mutex.RLock() // Use read lock for reading only
	defer bc.mutex.RUnlock()

	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) IsValid() bool {
	bc.mutex.RLock()
	defer bc.mutex.RUnlock()

	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

		if !currentBlock.IsValidHash() {
			return false
		}
	}
	return true
}

// GetPendingTransactions returns a copy of the pending transactions
func (bc *Blockchain) GetPendingTransactions() []Transaction {
	bc.mutex.RLock()
	defer bc.mutex.RUnlock()

	// Create a copy to avoid race conditions
	transactions := make([]Transaction, len(bc.PendingTransactions))
	copy(transactions, bc.PendingTransactions)

	return transactions
}

// AddMinedBlock adds a pre-mined block to the blockchain
func (bc *Blockchain) AddMinedBlock(block *Block) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	bc.Blocks = append(bc.Blocks, block)
}

// ClearPendingTransactions clears all pending transactions
func (bc *Blockchain) ClearPendingTransactions() {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	bc.PendingTransactions = []Transaction{}
}
