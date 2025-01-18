package db

import (
	"log"
	"time"
)

type Task struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Status     string    `json:"status"`
	Payload    string    `json:"payload"`
	RetryCount int       `json:"retry_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func AutoMigrate() {
	err := DB.AutoMigrate(&Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed!")
}
