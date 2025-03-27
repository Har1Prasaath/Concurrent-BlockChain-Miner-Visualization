package api

import (
	"blockchain-visualizer/blockchain"

	"github.com/gorilla/mux"
)

// SetupRoutesWithMining configures all the routes for our blockchain API
func SetupRoutesWithMining(router *mux.Router, bc *blockchain.Blockchain, numMiners int) {
	router.HandleFunc("/transactions/new", CreateTransactionHandler(bc)).Methods("POST")
	router.HandleFunc("/mine", MineBlockHandlerWithConcurrency(bc, numMiners)).Methods("GET")
	router.HandleFunc("/chain", GetBlockchainHandler(bc)).Methods("GET")
}

// Keep the original SetupRoutes for backward compatibility
func SetupRoutes(router *mux.Router, bc *blockchain.Blockchain) {
	SetupRoutesWithMining(router, bc, 1) // Default to 1 miner if not specified
}
