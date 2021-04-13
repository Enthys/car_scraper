package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := os.Getenv("DATABASE_DSN")

	db, err := attemptConnection(dsn, &ConnectionAttempt{
		MaxAttempts: 5,
		RetryDelay:  2 * time.Second,
	})

	if err != nil {
		panic("Failed to connecto to database")
	}

	DB = db
	log.Println("Connected to database!")
}

type ConnectionAttempt struct {
	MaxAttempts  int8
	RetryDelay   time.Duration
	attemptCount int8
}

func (attempt ConnectionAttempt) NewAttempt(maxAttempts int8, retryDelay time.Duration) *ConnectionAttempt {
	return &ConnectionAttempt{
		MaxAttempts: maxAttempts,
		RetryDelay:  retryDelay,
	}
}

func attemptConnection(dsn string, attempt *ConnectionAttempt) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		attempt.attemptCount += 1

		if attempt.attemptCount == attempt.MaxAttempts {
			log.Fatalf("Failed to connect")
			return nil, err
		}

		time.Sleep(attempt.RetryDelay)

		return attemptConnection(dsn, attempt)
	}

	return db, nil
}
