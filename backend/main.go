package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"blockchain-visualizer/api"
	"blockchain-visualizer/blockchain"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize the blockchain
	blockchain := blockchain.NewBlockchain()

	// Set up the router
	router := mux.NewRouter()

	// Set the number of miners based on available CPU cores
	numMiners := runtime.NumCPU()
	fmt.Printf("Using %d miners for concurrent mining with spanning tree termination\n", numMiners)

	// Define API routes with mining options
	api.SetupRoutesWithMining(router, blockchain, numMiners)

	// CORS configuration
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3000"}, // Allow all origins for testing
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
		Debug:            true, // Enable for debugging
	}
	corsHandler := cors.New(corsOptions)
	handler := corsHandler.Handler(router)

	// Start the server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
