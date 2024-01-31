package models

import "time"

const (
	AccountNoLength = 8
)

type Account struct {
	AccountNo string    `gorm:"primaryKey" json:"account_no"`
	UserID    uint32    `json:"user_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignKey:UserID"`
}
