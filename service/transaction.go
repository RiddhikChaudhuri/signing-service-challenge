package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"github.com/google/uuid"
)

func AddNewSignature(sigRequest *domain.SignatureRequest, s persistence.Storage) (signatureResponse *domain.SignatureResponse, err error) {
	log.Printf("Creating New Signature For The Device Id %s\n", sigRequest.DeviceID)
	signedDevice, _ := s.FindSignatureDeviceByID(sigRequest.DeviceID)
	var wg sync.WaitGroup
	if signedDevice == nil {
		return nil, errors.New("device key token pair not present in our system")
	}
	signResponse := &domain.SignatureResponse{}
	lastSignature := ""
	wg.Add(1)
	newCounterValue := signedDevice.SignatureCounter.Add(1)
	wg.Done()
	if newCounterValue == 1 {
		lastSignature = base64.StdEncoding.EncodeToString([]byte(sigRequest.DeviceID))
	} else {
		log.Printf("Current Counter Value %d", signedDevice.SignatureCounter.Load())
		signature, _ := s.FindCurrentSignatureByDeviceId(sigRequest.DeviceID)
		lastSignature = signature.SignatureBase64
	}
	securedDataToBeSigned := fmt.Sprintf("%d_%s_%s", newCounterValue, sigRequest.Data, lastSignature)
	var signer crypto.Signer
	// Now you need to sign securedDataToBeSigned using Signer.sign
	switch signedDevice.SignatureAlgorithm {
	case domain.RSA:
		marshaller := crypto.RSAMarshaler{}
		privateRsaKey, err := marshaller.Unmarshal(signedDevice.PrivateKey.Data)
		if err != nil {
			return nil, errors.New("Failed To Convert RSA Private Key From key bytes with error Details" + err.Error())
		}
		rsaSigner, err := crypto.NewRSASigner(privateRsaKey.Private)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		signer = rsaSigner

	case domain.ECC:
		marshaller := crypto.ECCMarshaler{}
		privateECCKey, err := marshaller.Decode(signedDevice.PrivateKey.Data)
		if err != nil {
			return nil, errors.New("Failed To Convert ECC Private Key From key bytes with error Details" + err.Error())
		}
		eccSigner, err := crypto.NewECCSigner(privateECCKey.Private)
		if err != nil {
			return nil, errors.New("failed to create ECC signer")
		}
		signer = eccSigner

	default:
		return nil, errors.New("invalid signature algorithm")
	}

	signatureBytes, err := signer.Sign([]byte(securedDataToBeSigned))
	if err != nil {
		return nil, errors.New(err.Error())
	}
	signatureBase64 := base64.StdEncoding.EncodeToString(signatureBytes)
	signatureEntity := ConvertToSignatureDomain(sigRequest, signatureBase64, securedDataToBeSigned)
	s.CreateSignTransaction(&signatureEntity)
	s.UpdateCounter(signedDevice.ID)
	wg.Wait() // Moved Wait here
	// Construct the SignatureResponse
	signResponse.Signature = signatureBase64
	signResponse.SignedData = securedDataToBeSigned
	return signResponse, nil
}

func ConvertToSignatureDomain(d *domain.SignatureRequest, signature string, signedData string) domain.Signature {
	return domain.Signature{
		ID:              uuid.New().String(),
		DeviceID:        d.DeviceID,
		SignatureBase64: signature,
		SignedData:      signedData,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}
}
