package persistence

import "github.com/fiskaly/coding-challenges/signing-service-challenge/domain"

type Storage interface {
	CreateSignatureDevice(signatureDevice *domain.SignatureDevice) (*domain.SignatureDevice, error)
	FindSignatureDeviceByID(deviceID string) (signedDevice *domain.SignatureDevice, err error)
	CreateSignature(sigRequest *domain.SignatureRequest) (signatureResponse *domain.SignatureResponse, err error)
}
