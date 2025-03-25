package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Transaction struct {
	ID        string
	Sender    string
	Recipient string
	Amount    float64
	Timestamp int64
}

func NewTransaction(sender, recipient string, amount float64) Transaction {
	tx := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
	}
	tx.ID = tx.CalculateHash()
	return tx
}

func (tx *Transaction) CalculateHash() string {
	record := fmt.Sprintf("%s%s%f%d", tx.Sender, tx.Recipient, tx.Amount, tx.Timestamp)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (tx *Transaction) ToString() string {
	return fmt.Sprintf("Transaction{ID: %s, Sender: %s, Recipient: %s, Amount: %.2f, Timestamp: %d}",
		tx.ID, tx.Sender, tx.Recipient, tx.Amount, tx.Timestamp)
}
