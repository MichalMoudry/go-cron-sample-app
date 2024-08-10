package service

import (
	"log"
	"time"

	"github.com/google/uuid"
)

// A simple job that prints a string.
func PrintJob(id uuid.UUID) {
	now := time.Now()
	calc := 1000 + 543*43/now.Second()
	log.Printf("Executed print job (%v).\n", calc)
	log.Println(id)
}
