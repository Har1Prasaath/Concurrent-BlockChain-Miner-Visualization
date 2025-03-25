package api

import (
	"encoding/json"
	"net/http"

	"blockchain-visualizer/blockchain"
)

type TransactionRequest struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

type BlockResponse struct {
	Message    string            `json:"message"`
	Block      *blockchain.Block `json:"block"`
	BlockIndex int               `json:"blockIndex"`
}

type TransactionResponse struct {
	Message string            `json:"message"`
	Block   *blockchain.Block `json:"block"`
}

type BlockchainResponse struct {
	Chain  []*blockchain.Block `json:"chain"`
	Length int                 `json:"length"`
}

func CreateTransactionHandler(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req TransactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		transaction := blockchain.NewTransaction(req.Sender, req.Recipient, req.Amount)

		// Add to pending transactions (usually would be stored then mined)
		newBlock := bc.AddBlock([]blockchain.Transaction{transaction})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TransactionResponse{
			Message: "Transaction added successfully",
			Block:   newBlock,
		})
	}
}

func MineBlockHandler(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// In a real implementation, this would mine pending transactions
		// For this example, we'll just create a new block with a dummy transaction
		dummyTx := blockchain.NewTransaction("system", "miner", 1.0)
		newBlock := bc.AddBlock([]blockchain.Transaction{dummyTx})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(BlockResponse{
			Message:    "New block mined",
			BlockIndex: newBlock.Index,
			Block:      newBlock,
		})
	}
}

func GetBlockchainHandler(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := BlockchainResponse{
			Chain:  bc.Blocks,
			Length: len(bc.Blocks),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
