package services

import (
	"github.com/nguyenaivuong/litebank/internal/models"
	"gorm.io/gorm"
)

// EntryService provides methods for account-related operations
type EntryService struct {
	db *gorm.DB
}

// NewEntryService creates a new instance of EntryService
func NewEntryService(db *gorm.DB) *EntryService {
	return &EntryService{db: db}
}

// CreateEntry creates a new entry using GORM
func (s *EntryService) CreateEntry(entry *models.Entry) error {
	return s.db.Create(entry).Error
}

// GetEntry retrieves an entry by ID using GORM
func (s *EntryService) GetEntry(id int64) (entry *models.Entry, err error) {
	err = s.db.First(&entry, "id = ?", id).Error
	return entry, err
}
