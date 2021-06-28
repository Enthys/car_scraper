package main

import (
	"car_scraper/cmd"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if pathExists(dir + "/.env") {
		err = godotenv.Overload(dir + "/.env")
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
