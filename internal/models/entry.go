package models

import "time"

type Entry struct {
	ID            int64  `json:"id"`
	AccountNumber string `json:"account_id"`
	// can be negative or positive
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
