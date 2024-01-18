package service

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"github.com/google/uuid"
)

func AddNewSignatureDevice(csdr *domain.CreateSignatureDeviceRequest, s persistence.Storage) (*domain.CreateSignatureDeviceResponse, error) {
	// Convert the request to a SignatureDevice
	log.Println("Converting Add Device Request To Domain")
	signatureDevice := ConvertToDomain(csdr)
	pub, private, _ := crypto.GetKeyPair(signatureDevice.SignatureAlgorithm)

	signatureDevice.PublicKey = pub
	signatureDevice.PrivateKey = domain.EncryptedKey{Algorithm: signatureDevice.SignatureAlgorithm, Data: private}

	storageResult, err := s.CreateSignatureDevice(&signatureDevice)
	if err != nil {
		return nil, err
	}

	// Return the created signature device response
	response := &domain.CreateSignatureDeviceResponse{
		ID:    storageResult.ID,
		Label: storageResult.Label,
	}

	return response, nil
}

func ConvertToDomain(d *domain.CreateSignatureDeviceRequest) domain.SignatureDevice {
	return domain.SignatureDevice{
		ID:                 uuid.New().String(),
		SignatureAlgorithm: getAlgorithm(d.Algorithm),
		Label:              d.Label,
		SignatureCounter:   atomic.Int64{},
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

func getAlgorithm(algorithm string) string {
	switch algorithm {
	case "RSA":
		return domain.RSA
	case "ECC":
		return domain.ECC
	default:
		return domain.ECC
	}
}
