package services

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nguyenaivuong/litebank/internal/models"
	"github.com/nguyenaivuong/litebank/internal/utils"
	"gorm.io/gorm"
)

type SignupService interface {
	Signup(ctx *gin.Context, user *models.User) error
}

type signupService struct {
	db *gorm.DB
}

func NewSignupService(db *gorm.DB) SignupService {
	return &signupService{db: db}
}

/*
func (s *signupService) Signup(ctx *gin.Context, user *models.User) error {
	// Check if username already exists
	existingUser, err := s.existUser(user.Username)
	if err != nil {
		utils.LogWithCallerInfo("")
		return err
	}

	if existingUser != nil {
		err := errors.New("username already exists")
		utils.LogWithCallerInfo(err.Error())
		return err
	}

	// Save user to database
	err = s.db.Create(&user).Error
	if err != nil {
		utils.LogWithCallerInfo(err.Error())
		return err
	}

	// Create account for user
	account := &models.Account{
		UserID:    user.ID,
		Balance:   0,
		AccountNo: utils.GenerateAccountNumber(models.AccountNoLength),
		CreatedAt: time.Now(),
	}

	if err := s.db.Create(account).Error; err != nil {
		utils.LogWithCallerInfo(err.Error())
		return err
	}

	return nil
}
*/

func (s *signupService) Signup(ctx *gin.Context, user *models.User) error {
	// Start a new transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		utils.LogWithCallerInfo(tx.Error.Error())
		return tx.Error
	}

	// Use defer to ensure the transaction is rolled back if an error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if username already exists
	existingUser, err := s.existUser(user.Username)
	if err != nil {
		utils.LogWithCallerInfo(err.Error())
		tx.Rollback()
		return err
	}

	if existingUser != nil {
		err := errors.New("username already exists")
		utils.LogWithCallerInfo(err.Error())
		tx.Rollback()
		return err
	}

	// Save user to database
	err = tx.Create(&user).Error
	if err != nil {
		utils.LogWithCallerInfo(err.Error())
		tx.Rollback()
		return err
	}

	// Create account for user
	account := &models.Account{
		UserID:    user.ID,
		Balance:   0,
		AccountNo: utils.GenerateAccountNumber(models.AccountNoLength),
		CreatedAt: time.Now(),
	}

	if err := tx.Create(account).Error; err != nil {
		utils.LogWithCallerInfo(err.Error())
		tx.Rollback()
		return err
	}

	// If everything went well, commit the transaction
	if err := tx.Commit().Error; err != nil {
		utils.LogWithCallerInfo(err.Error())
		return err
	}

	return nil
}

func (s *signupService) existUser(username string) (*models.User, error) {
	var user *models.User
	if err := s.db.First(&user, "username", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.LogWithCallerInfo(err.Error())
		return nil, err
	}

	return user, nil
}
