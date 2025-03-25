package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        int           `json:"Index"`
	Timestamp    int64         `json:"Timestamp"`
	Transactions []Transaction `json:"Transactions"`
	PreviousHash string        `json:"PreviousHash"`
	Hash         string        `json:"Hash"`
	Nonce        int           `json:"Nonce"`
	Difficulty   int           `json:"Difficulty"`
}

func NewBlock(index int, previousHash string, transactions []Transaction) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PreviousHash: previousHash,
		Nonce:        0,
		Difficulty:   4, // Difficulty can be adjusted
	}
	block.MineBlock()
	return block
}

func (b *Block) CalculateHash() string {
	data := strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp, 10) + b.PreviousHash + strconv.Itoa(b.Nonce)
	for _, tx := range b.Transactions {
		data += tx.ID
	}
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (b *Block) IsValidHash() bool {
	return strings.HasPrefix(b.Hash, strings.Repeat("0", b.Difficulty))
}

func (b *Block) MineBlock() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Hash = b.CalculateHash()
		if strings.HasPrefix(b.Hash, target) {
			break
		}
		b.Nonce++
	}
}
