package api

import (
	"blockchain-visualizer/blockchain"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all the routes for our blockchain API
func SetupRoutes(router *mux.Router, bc *blockchain.Blockchain) {
	router.HandleFunc("/transactions/new", CreateTransactionHandler(bc)).Methods("POST")
	router.HandleFunc("/mine", MineBlockHandler(bc)).Methods("GET")
	router.HandleFunc("/chain", GetBlockchainHandler(bc)).Methods("GET")
}
