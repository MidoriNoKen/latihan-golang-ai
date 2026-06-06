package main

import (
	"log"
	"os"

	"github.com/MidoriNoKen/latihan-golang-ai/config"
	"github.com/MidoriNoKen/latihan-golang-ai/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Using system environment variables.")
	}

	// 1. Initialize Database Connection
	config.InitDB()

	// 2. Setup Router
	r := routes.SetupRouter()

	// 3. Get server port from env variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)

	// 4. Start Server
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
