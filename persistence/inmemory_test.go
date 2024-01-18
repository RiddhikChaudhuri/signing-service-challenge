package persistence

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

func TestInMemoryStorage(t *testing.T) {
	// Create a new instance of InMemoryStorage for testing
	store := NewMemoryStore()

	// Create a test signature device
	deviceID := uuid.NewString()
	signatureDevice := &domain.SignatureDevice{
		ID:                 deviceID,
		Label:              "TestDevice",
		SignatureAlgorithm: domain.RSA,
		PrivateKey:         domain.EncryptedKey{Algorithm: "RSA", Data: []byte("test_private_key")},
		PublicKey:          []byte("test_public_key"),
		SignatureCounter:   atomic.Int64{},
		CreatedAt:          time.Now().UTC(),
		UpdatedAt:          time.Now().UTC(),
	}

	// Test CreateSignatureDevice
	createdDevice, err := store.CreateSignatureDevice(signatureDevice)
	if err != nil {
		t.Errorf("Error Expected To Be Nil But Found %v", err)
	}
	if createdDevice == nil {
		t.Errorf("Created Device Expected To Be Not Nil But is Nil ")
	}
	if createdDevice != signatureDevice {
		t.Errorf("Created Device Expected To Be Equal To Signature Device But Is Not ")
	}

	// Test FindSignatureDeviceByID
	foundDevice, err := store.FindSignatureDeviceByID(deviceID)
	if err != nil {
		t.Errorf("Error Expected To Be Nil But Found %v", err)
	}
	if foundDevice == nil {
		t.Errorf("Found Device Expected To Be Not Nil But is Nil ")
	}
	if createdDevice != foundDevice {
		t.Errorf("Created Device Expected To Be Equal To Found Device But Is Not ")
	}

	// Test FindSignatureDeviceByID for non-existent device
	nonExistentDeviceID := uuid.NewString()
	_, err = store.FindSignatureDeviceByID(nonExistentDeviceID)
	if err == nil {
		t.Errorf("Error Expected To Be Found But was Nil %v", err)
	}
	if err.Error() != "device not found" {
		t.Errorf("Error Expected To Be Device Not Found But was %v", err.Error())
	}

	// Add more test cases for other methods...

	// Cleanup (if needed)

	// Remove the test device from the storage
	delete(store.SignatureDevices, deviceID)
}

// Add more tests for other methods if needed...
