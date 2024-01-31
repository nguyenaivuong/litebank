package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenaivuong/litebank/internal/models"
	"gorm.io/gorm"
)

// AccountService provides methods for account-related operations
type AccountService struct {
	db *gorm.DB
}

// NewAccountService creates a new instance of AccountService
func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{db: db}
}

// GetAccountByID retrieves an account by ID
func (s *AccountService) GetAccountByID(accountID string) (*models.Account, error) {
	var account models.Account
	err := s.db.First(&account, "account_number = ?", accountID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Handle not found gracefully
		}
		return nil, err
	}
	return &account, nil
}

// UpdateAccount updates an existing account
func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.db.Save(account).Error
}

// DeleteAccount deletes an account by its primary key
func (s *AccountService) DeleteAccount(accountID string) error {
	return s.db.Delete(&models.Account{}, "account_number = ?", accountID).Error
}

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	UserID uint32 `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

// ListAccounts fetches a list of all accounts
func (s *AccountService) ListAccounts(arg ListAccountsParams) ([]*models.Account, error) {
	var accounts []*models.Account
	err := s.db.Preload("User").Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccountByUserID fetches an account associated with a given user ID
func (s *AccountService) GetAccountByUserID(userID uint32) (*models.Account, error) {
	var account models.Account
	err := s.db.Preload("User").First(&account, "user_id = ?", userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Handle not found gracefully
		}
		return nil, err
	}
	return &account, nil
}

func (s *AccountService) ValidAccount(ctx *gin.Context, accountID string, currency string) (*models.Account, bool) {
	account, err := s.GetAccountByID(accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return account, false
	}

	if account == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return account, false
	}

	return account, true
}
