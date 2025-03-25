package main

import (
	"fmt"
	"log"
	"net/http"

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

	// Define API routes
	api.SetupRoutes(router, blockchain)

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
