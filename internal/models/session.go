package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID       uint32
	RefreshToken string
	UserAgent    string
	ClientIP     string `gorm:"column:client_ip"`
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
