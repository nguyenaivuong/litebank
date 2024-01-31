package services

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/nguyenaivuong/litebank/models"
// 	"gorm.io/gorm"
// )

// type TransactionService interface {
// 	Deposit(ctx context.Context, accountID int, amount float64) error
// 	Withdraw(ctx context.Context, accountID int, amount float64) error
// }

// type transactionService struct {
// 	db *gorm.DB
// }

// func NewTransactionService(db *gorm.DB) TransactionService {
// 	return &transactionService{db: db}
// }

// func (s *transactionService) Deposit(ctx context.Context, accountID int, amount float64) error {
// 	// Check if account exists
// 	var account *models.Account
// 	err := s.db.First(&account, accountID).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Update account balance
// 	account.Balance += amount
// 	err = s.db.Save(&account).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Save transaction record
// 	transaction := &models.Transaction{
// 		AccountID: accountID,
// 		Type:      "deposit",
// 		Amount:    amount,
// 		Timestamp: time.Now(),
// 	}
// 	err = s.db.Create(transaction).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *transactionService) Withdraw(ctx context.Context, accountID int, amount float64) error {
// 	// Check if account exists
// 	var account *models.Account
// 	err := s.db.First(&account, accountID).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Validate sufficient balance
// 	if account.Balance < amount {
// 		return errors.New("insufficient funds")
// 	}

// 	// Update account balance
// 	account.Balance -= amount
// 	err = s.db.Save(&account).Error
// 	if err != nil {
// 		return err
// 	}

// 	// Save transaction record
// 	transaction := &models.Transaction{
// 		AccountID: accountID,
// 		Type:      "withdraw",
// 		Amount:    amount,
// 		Timestamp: time.Now(),
// 	}
// 	err = s.db.Create(transaction).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
