package main

import (
	"car_scraper/cmd"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if pathExists(".env") {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return
		}
	}

	cmd.Execute()
}

func pathExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}