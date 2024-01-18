package persistence

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type Storage interface {
	CreateSignatureDevice(signatureDevice *domain.SignatureDevice) (*domain.SignatureDevice, error)
	FindSignatureDeviceByID(deviceID string) (signedDevice *domain.SignatureDevice, err error)
	CreateSignTransaction(sigRequest *domain.Signature) (signatureResponse *domain.Signature, err error)
	FindCurrentSignatureByDeviceId(deviceID string) (signature *domain.Signature, err error)
	UpdateCounter(deviceID string) error
	FindAllSignatureDevices() (signedDevices []*domain.SignatureDevice)
	FindAllSignatures() (signedTransactions []*domain.Signature)
}
