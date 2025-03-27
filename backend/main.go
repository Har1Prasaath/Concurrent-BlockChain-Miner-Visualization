package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"blockchain-visualizer/api"
	"blockchain-visualizer/blockchain"
	"blockchain-visualizer/miner"

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

	// Initialize deadlock detector
	detector := miner.NewDeadlockDetector()

	// Create a sample deadlock for demonstration purposes
	// fmt.Println("▸▸▸ Creating sample deadlock scenario for demonstration")
	// detector.AddAllocation(1, 1)
	// detector.AddWaitFor(1, 2)
	// detector.AddAllocation(2, 2)
	// detector.AddWaitFor(2, 1)

	// Run deadlock detection immediately at startup
	fmt.Println("\n▸▸▸ Running initial deadlock detection...")
	detector.PrintDeadlocks()

	// Create a stop channel for graceful shutdown
	stopChan := make(chan struct{})

	// Start a goroutine that periodically checks for deadlocks every 5 seconds
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("\n▸▸▸ Running scheduled deadlock detection...")
				detector.PrintDeadlocks()
			case <-stopChan:
				return
			}
		}
	}()

	// CORS configuration
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3000"}, // Allow all origins for testing
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
		Debug:            false, // Disable for cleaner logs
	}
	corsHandler := cors.New(corsOptions)
	handler := corsHandler.Handler(router)

	// This call blocks until the server is shut down
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

	// These lines will only execute if the server shuts down gracefully
	close(stopChan)
}
