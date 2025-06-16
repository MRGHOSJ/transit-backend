package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"transit-backend/internal/router"
	"transit-backend/internal/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file (using system env vars instead): %v", err)
	}

	data, err := utils.LoadTransportData("data/transport.json")
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	r := router.Setup(data)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Println("🚀 Server running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
