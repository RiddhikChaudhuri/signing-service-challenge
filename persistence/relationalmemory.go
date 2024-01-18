package persistence

import (
	"errors"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	DB *gorm.DB
}

// NewPostgresStorage creates a new instance of PostgresStorage.
func NewPostgresStorage(db *gorm.DB) *PostgresStorage {
	return &PostgresStorage{
		DB: db,
	}
}

// CreateSignatureDevice creates a new signature device in the PostgreSQL storage.
func (s *PostgresStorage) CreateSignatureDevice(signatureDevice *domain.SignatureDevice) error {
	result := s.DB.Create(signatureDevice)
	return result.Error
}

// FindSignatureDeviceByID finds a signature device by ID in the PostgreSQL storage.
func (s *PostgresStorage) FindSignatureDeviceByID(deviceID string) (*domain.SignatureDevice, error) {
	var device domain.SignatureDevice
	result := s.DB.First(&device, "id = ?", deviceID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("device not found")
		}
		return nil, result.Error
	}

	return &device, nil
}
