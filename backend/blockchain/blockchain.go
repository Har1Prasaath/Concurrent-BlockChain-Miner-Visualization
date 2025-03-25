package blockchain

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(0, "", []Transaction{})
	bc := &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
	return bc
}

func (bc *Blockchain) AddBlock(transactions []Transaction) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, prevBlock.Hash, transactions)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

// GetLatestBlock returns the most recently added block
func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

// IsValid checks if the blockchain is valid
func (bc *Blockchain) IsValid() bool {
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
