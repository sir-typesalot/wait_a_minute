package util

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func VaultGet(key string) (secret string) {
	// Load in .env file
	err := godotenv.Load()
	// Error handling
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Read in value and return
	secret = os.Getenv(key)
	return
}
