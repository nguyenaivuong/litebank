package services

import (
	"errors"

	"github.com/nguyenaivuong/litebank/internal/models"
	"gorm.io/gorm"
)

type LoginService interface {
	Login(username, password string) (string, error)
}

type loginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) LoginService {
	return &loginService{db: db}
}

func (s *loginService) Login(username, password string) (string, error) {
	// Fetch user by username
	var user models.User
	err := s.db.First(&user, "username = ?", username).Error
	if err != nil {
		return "", err
	}

	// Verify password
	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := createToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	// Verify password against stored hash
	// This is just a placeholder
	return password == hash
}

func createToken(userID uint32) (string, error) {
	// Implement JWT token creation logic
	// This is just a placeholder
	return "dummy-token", nil
}
