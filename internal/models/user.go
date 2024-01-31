package models

import (
	"time"
)

type User struct {
	ID                uint32    `gorm:"primaryKey" json:"id"`
	Username          string    `json:"username" gorm:"uniqueIndex;not null"`
	Password          string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	Accounts          *Account  `gorm:"foreignKey:UserID"`
}
