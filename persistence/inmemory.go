package persistence

import (
	"errors"
	"log"
	"sync"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type InMemoryStorage struct {
	SignatureDevices map[string]*domain.SignatureDevice
	mutex            sync.Mutex
}

func NewMemoryStore() *InMemoryStorage {
	m := InMemoryStorage{
		SignatureDevices: make(map[string]*domain.SignatureDevice),
	}
	return &m
}

// CreateSignatureDevice creates a new signature device in the in-memory storage.
func (s *InMemoryStorage) CreateSignatureDevice(signatureDevice *domain.SignatureDevice) (signedDevice *domain.SignatureDevice, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Store the signature device in memory
	s.SignatureDevices[signatureDevice.ID] = signatureDevice
	log.Printf("Printing The stored Device %v\n", s.SignatureDevices[signatureDevice.ID])
	return signatureDevice, nil
}

func (s *InMemoryStorage) FindSignatureDeviceByID(deviceID string) (signedDevice *domain.SignatureDevice, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Printf("Looking Into In-Memory Db With Device Id %v", deviceID)
	// Retrieve the signature device from memory
	device, found := s.SignatureDevices[deviceID]
	log.Printf("Printing The stored Device %v\n", device)
	if found {
		// Create a new pointer to the device and return it
		result := *device
		return &result, nil
	}

	return nil, errors.New("device not found")
}

func (s *InMemoryStorage) CreateSignature(sigRequest *domain.SignatureRequest) (signatureResponse *domain.SignatureResponse, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("Creating New Signature For The Device Id %s\n", sigRequest.DeviceID)

	return nil, nil
}
