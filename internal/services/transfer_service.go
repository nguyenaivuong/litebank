package services

import (
	"github.com/nguyenaivuong/litebank/internal/models"
	"gorm.io/gorm"
)

// TransferService provides methods for account-related operations
type TransferService struct {
	db *gorm.DB
}

// NewTransferService creates a new instance of TransferService
func NewTransferService(db *gorm.DB) *TransferService {
	return &TransferService{db: db}
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount models.Account  `json:"from_account"`
	ToAccount   models.Account  `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

func (s *TransferService) CreateTransfer(transfer *models.Transfer) error {
	return s.db.Create(&transfer).Error
}

func (s *TransferService) GetTransferByID(db *gorm.DB, id int64) (transfer *models.Transfer, err error) {
	err = s.db.First(&transfer, id).Error
	return transfer, err
}

func (s *TransferService) ListTransfers() (transfers []*models.Transfer, err error) {
	err = s.db.Find(&transfers).Error
	return transfers, err
}

func (s *TransferService) TransferTx(arg models.TransferTxParams) (result TransferTxResult, err error) {
	// Check if account exists
	var fromAccount *models.Account
	err = s.db.First(&fromAccount, "account_number = ?", arg.FromAccountNumber).Error
	if err != nil {
		return result, err
	}

	var toAccount *models.Account
	err = s.db.First(&toAccount, "account_number = ?", arg.ToAccountNumber).Error
	if err != nil {
		return result, err
	}

	// Check if account has sufficient balance
	if fromAccount.Balance < arg.Amount {
		return result, err
	}

	// Update account balance
	fromAccount.Balance -= arg.Amount
	toAccount.Balance += arg.Amount

	// Create entries
	fromEntry := models.Entry{
		AccountNumber: arg.FromAccountNumber,
		Amount:        -arg.Amount,
	}
	toEntry := models.Entry{
		AccountNumber: arg.ToAccountNumber,
		Amount:        arg.Amount,
	}

	// Create transfer
	transfer := models.Transfer{
		FromAccountNumber: arg.FromAccountNumber,
		ToAccountNumber:   arg.ToAccountNumber,
		Amount:            arg.Amount,
	}

	// Create database transaction
	tx := s.db.Begin()
	err = tx.Create(&fromEntry).Error
	if err != nil {
		tx.Rollback()
		return result, err
	}
	err = tx.Create(&toEntry).Error
	if err != nil {
		tx.Rollback()
		return result, err
	}
	err = tx.Create(&transfer).Error
	if err != nil {
		tx.Rollback()
		return result, err
	}
	err = tx.Save(&fromAccount).Error
	if err != nil {
		tx.Rollback()
		return result, err
	}
	err = tx.Save(&toAccount).Error
	if err != nil {
		tx.Rollback()
		return result, err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return result, err
	}

	// Populate result
	result.Transfer = transfer
	result.FromAccount = *fromAccount
	result.ToAccount = *toAccount
	result.FromEntry = fromEntry
	result.ToEntry = toEntry

	return result, nil
}
