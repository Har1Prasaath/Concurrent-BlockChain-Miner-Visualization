package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Wallet represents a user's wallet
type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// NewWallet creates a new wallet
func NewWallet() Wallet {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return Wallet{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}
}

// SignTransaction signs a transaction
func (w Wallet) SignTransaction(tx Transaction) string {
	txData := fmt.Sprintf("%s%s%d", tx.Sender, tx.Receiver, tx.Amount)
	hash := sha256.Sum256([]byte(txData))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, w.PrivateKey, 0, hash[:])
	return hex.EncodeToString(signature)
}

// VerifyTransaction verifies a transaction's signature
func VerifyTransaction(tx Transaction, signature string, publicKey *rsa.PublicKey) bool {
	txData := fmt.Sprintf("%s%s%d", tx.Sender, tx.Receiver, tx.Amount)
	hash := sha256.Sum256([]byte(txData))
	sigBytes, _ := hex.DecodeString(signature)
	err := rsa.VerifyPKCS1v15(publicKey, 0, hash[:], sigBytes)
	return err == nil
}

// Transaction represents a transaction
type Transaction struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   int    `json:"amount"`
}

// NewTransaction creates a new transaction
func NewTransaction(sender, receiver string, amount int) Transaction {
	return Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}