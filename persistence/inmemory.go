package persistence

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type InMemoryStorage struct {
	SignatureDevices      map[string]*domain.SignatureDevice
	SignatureTransactions map[string]*domain.Signature
	mutex                 sync.Mutex
}

func NewMemoryStore() *InMemoryStorage {
	m := InMemoryStorage{
		SignatureDevices:      make(map[string]*domain.SignatureDevice),
		SignatureTransactions: make(map[string]*domain.Signature),
	}
	return &m
}

// CreateSignatureDevice creates a new signature device in the in-memory storage.
func (s *InMemoryStorage) CreateSignatureDevice(signatureDevice *domain.SignatureDevice) (signedDevice *domain.SignatureDevice, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Store the signature device in memory
	s.SignatureDevices[signatureDevice.ID] = signatureDevice
	return signatureDevice, nil
}

func (s *InMemoryStorage) FindSignatureDeviceByID(deviceID string) (signedDevice *domain.SignatureDevice, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Printf("Looking Into In-Memory Db With Device Id %v", deviceID)
	// Retrieve the signature device from memory
	device, found := s.SignatureDevices[deviceID]
	if found {
		return device, nil
	}

	return nil, errors.New("device not found")
}

func (s *InMemoryStorage) CreateSignTransaction(sigRequest *domain.Signature) (signatureResponse *domain.Signature, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.SignatureTransactions[sigRequest.DeviceID] = sigRequest
	return sigRequest, nil
}

func (s *InMemoryStorage) FindCurrentSignatureByDeviceId(deviceID string) (*domain.Signature, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Printf("Looking Into In-Memory Db For Signature With Device Id %v", deviceID)
	// Retrieve the signature device from memory
	signature, found := s.SignatureTransactions[deviceID]
	if found {
		// Create a new instance of the device and return it
		result := domain.Signature{
			ID:              signature.ID,
			DeviceID:        signature.DeviceID,
			SignatureBase64: signature.SignatureBase64,
			SignedData:      signature.SignedData,
		}
		return &result, nil
	}

	return nil, errors.New("signature not found")
}

func (s *InMemoryStorage) UpdateCounter(deviceID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if value, ok := s.SignatureDevices[deviceID]; ok {
		// Update other fields if needed
		s.SignatureDevices[deviceID] = value
		return nil
	}

	return fmt.Errorf("device with ID: %v not found", deviceID)
}

func (s *InMemoryStorage) FindAllSignatureDevices() (signedDevices []*domain.SignatureDevice) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Initialize the slice to hold the results
	signedDevices = make([]*domain.SignatureDevice, 0)

	// Iterate over the map and append each device to the slice
	for _, device := range s.SignatureDevices {
		signedDevices = append(signedDevices, device)
	}

	return signedDevices
}

func (s *InMemoryStorage) FindAllSignatures() (signedTransactions []*domain.Signature) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Initialize the slice to hold the results
	signedTransactions = make([]*domain.Signature, 0)

	// Iterate over the map and append each device to the slice
	for _, sigTransaction := range s.SignatureTransactions {
		signedTransactions = append(signedTransactions, sigTransaction)
	}

	return signedTransactions
}
