package services

import (
	"github.com/google/uuid"
	"github.com/nguyenaivuong/litebank/internal/models"
	"gorm.io/gorm"
)

// SessionService provides methods for session-related operations
type SessionService struct {
	db *gorm.DB
}

// NewSessionService creates a new instance of SessionService
func NewSessionService(db *gorm.DB) *SessionService {
	return &SessionService{db: db}
}

// Createsession creates a new session
func (s *SessionService) CreateSession(session *models.Session) error {
	return s.db.Create(session).Error
}

// GetsessionByID retrieves an session by ID
func (s *SessionService) GetSession(uuid uuid.UUID) (*models.Session, error) {
	var session *models.Session
	err := s.db.First(&session, "id = ?", uuid).Error
	return session, err
}
