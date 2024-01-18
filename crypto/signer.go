package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

// RSASigner represents an RSA signer.
type RSASigner struct {
	privateKey *rsa.PrivateKey
}

// NewRSASigner creates a new RSASigner instance.
func NewRSASigner(privateKey crypto.PrivateKey) (*RSASigner, error) {
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("invalid private key type, expected *rsa.PrivateKey")
	}
	return &RSASigner{privateKey: rsaPrivateKey}, nil
}

// Sign signs the provided data using RSA.
func (s *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, dataToBeSigned)
}

// ECCSigner represents an ECC signer.
type ECCSigner struct {
	privateKey *ecdsa.PrivateKey
}

// NewECCSigner creates a new ECCSigner instance.
func NewECCSigner(privateKey crypto.PrivateKey) (*ECCSigner, error) {
	ecdsaPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("invalid private key type, expected *ecdsa.PrivateKey")
	}
	return &ECCSigner{privateKey: ecdsaPrivateKey}, nil
}

func (signer *ECCSigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, signer.privateKey, dataToBeSigned)
	if err != nil {
		return nil, err
	}

	// print the signature
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	return signature, nil
}

func GetKeyPair(algorithm string) (public, private []byte, err error) {
	switch algorithm {
	case domain.RSA:
		rsaGenerator := RSAGenerator{}
		k, errGenerator := rsaGenerator.Generate()
		if errGenerator != nil {
			err = errGenerator
			return
		}

		marshaller := RSAMarshaler{}
		return marshaller.Marshal(*k)
	case domain.ECC:
		eccGenerator := ECCGenerator{}
		k, errGenerator := eccGenerator.Generate()
		if errGenerator != nil {
			err = errGenerator
			return
		}

		marshaller := ECCMarshaler{}
		return marshaller.Encode(*k)
	default:
		err = fmt.Errorf("wrong algorithm Type %s", algorithm)
		return
	}
}
