package models

import (
	"time"
)

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

type TransferTxParams struct {
	FromAccountNumber string  `json:"from_account_id"`
	ToAccountNumber   string  `json:"to_account_id"`
	Amount            float64 `json:"amount"`
}

type Transfer struct {
	ID                int64  `json:"id"`
	FromAccountNumber string `json:"from_account_id"`
	ToAccountNumber   string `json:"to_account_id"`
	// must be positive
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
