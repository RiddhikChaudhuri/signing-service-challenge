package service

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

func TestAddNewSignature(t *testing.T) {
	// Create a mock storage implementation for testing
	mockStorage := &MockStorage{}
	deviceID := "1c805988-d18c-4923-a334-003dbf010653"
	sampleEncodedData := "Hello!!How are you"

	// Create a sample signature request
	signatureRequest := &domain.SignatureRequest{
		DeviceID: deviceID,
		Data:     sampleEncodedData,
	}

	// Call the function with the mock storage
	signatureResponse, err := AddNewSignature(signatureRequest, mockStorage)

	// Assert the results
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if signatureResponse == nil {
		t.Error("Expected signatureResponse to be not nil, but got nil")
	}
}

// MockStorage is a mock implementation of the persistence.Storage interface for testing.
type MockStorage struct{}

func (m *MockStorage) CreateSignatureDevice(signatureDevice *domain.SignatureDevice) (*domain.SignatureDevice, error) {
	return nil, nil
}

func (m *MockStorage) FindSignatureDeviceByID(deviceID string) (*domain.SignatureDevice, error) {
	return &domain.SignatureDevice{
		ID:                 deviceID,
		Label:              "Unit-Test",
		SignatureAlgorithm: domain.RSA,
		PrivateKey: domain.EncryptedKey{
			Algorithm: domain.RSA,
			Data:      []byte("LS0tLS1CRUdJTiBQVUJMSUNfS0VZLS0tLS0KTUhZd0VBWUhLb1pJemowQ0FRWUZLNEVFQUNJRFlnQUUyeVZEajJwalhuS1dlSyt3Q1RDaGVvc2h5RXd1MGR6cApEUVZaajA4MWNuQVVWbUtvZk1ZQ3MrQng4VVNoTkxhVHUxcWl3YkY3NThGb1puYytuek43ZStpbmZORzBWR0hMCm5iU3JKdURteVZ6bU02bjg0L1Zqb3hvaG5DRTJ5UmIwCi0tLS0tRU5EIFBVQkxJQ19LRVktLS0tLQo="),
		},
		PublicKey:        []byte("LS0tLS1CRUdJTiBQVUJMSUNfS0VZLS0tLS0KTUhZd0VBWUhLb1pJemowQ0FRWUZLNEVFQUNJRFlnQUUyeVZEajJwalhuS1dlSyt3Q1RDaGVvc2h5RXd1MGR6cApEUVZaajA4MWNuQVVWbUtvZk1ZQ3MrQng4VVNoTkxhVHUxcWl3YkY3NThGb1puYytuek43ZStpbmZORzBWR0hMCm5iU3JKdURteVZ6bU02bjg0L1Zqb3hvaG5DRTJ5UmIwCi0tLS0tRU5EIFBVQkxJQ19LRVktLS0tLQo="),
		SignatureCounter: atomic.Int64{},
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}, nil
}

func (m *MockStorage) FindCurrentSignatureByDeviceId(deviceID string) (*domain.Signature, error) {
	return &domain.Signature{
		ID:              uuid.New().String(),
		DeviceID:        "1c805988-d18c-4923-a334-003dbf010653",
		SignatureBase64: "fFQ0AJROCfZe9pvUwJbgOnvf1WgTAoTp6MJNi2jXK89OlZFnd3rAKMb/a+EQEQiYxa7xHQ8iONfOzeI+NHXoQ6qbK2SO3iMBGljCF1hFPm2L3fLttVTd5fOhFrHfyXmB",
		SignedData:      "1_Hello!!How are you_MWM4MDU5ODgtZDE4Yy00OTIzLWEzMzQtMDAzZGJmMDEwNjUz",
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}, nil
}

func (m *MockStorage) UpdateCounter(deviceID string) error {
	return nil
}

func (m *MockStorage) FindAllSignatureDevices() []*domain.SignatureDevice {
	return []*domain.SignatureDevice{{
		ID:                 "1c805988-d18c-4923-a334-003dbf010653",
		Label:              "Unit-Test",
		SignatureAlgorithm: domain.RSA,
		PrivateKey: domain.EncryptedKey{
			Algorithm: domain.RSA,
			Data:      []byte("test_private_key"),
		},
		PublicKey:        []byte("test_public_key"),
		SignatureCounter: atomic.Int64{},
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	},
	}
}

func (m *MockStorage) FindAllSignatures() []*domain.Signature {
	return []*domain.Signature{
		{
			ID:              uuid.New().String(),
			DeviceID:        "1c805988-d18c-4923-a334-003dbf010653",
			SignatureBase64: "fFQ0AJROCfZe9pvUwJbgOnvf1WgTAoTp6MJNi2jXK89OlZFnd3rAKMb/a+EQEQiYxa7xHQ8iONfOzeI+NHXoQ6qbK2SO3iMBGljCF1hFPm2L3fLttVTd5fOhFrHfyXmB",
			SignedData:      "1_Hello!!How are you_MWM4MDU5ODgtZDE4Yy00OTIzLWEzMzQtMDAzZGJmMDEwNjUz",
			CreatedAt:       time.Now().UTC(),
			UpdatedAt:       time.Now().UTC(),
		},
	}
}

func (m *MockStorage) CreateSignTransaction(signature *domain.Signature) (*domain.Signature, error) {
	return &domain.Signature{
		ID:              uuid.New().String(),
		DeviceID:        signature.DeviceID,
		SignatureBase64: "fFQ0AJROCfZe9pvUwJbgOnvf1WgTAoTp6MJNi2jXK89OlZFnd3rAKMb/a+EQEQiYxa7xHQ8iONfOzeI+NHXoQ6qbK2SO3iMBGljCF1hFPm2L3fLttVTd5fOhFrHfyXmB",
		SignedData:      "1_Hello!!How are you_MWM4MDU5ODgtZDE4Yy00OTIzLWEzMzQtMDAzZGJmMDEwNjUz",
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}, nil
}
