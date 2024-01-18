package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

// RSASigner represents an RSA signer.
type RSASigner struct {
	privateKey *rsa.PrivateKey
}

// NewRSASigner creates a new RSASigner instance.
func NewRSASigner(privateKey crypto.PrivateKey) (*RSASigner, error) {
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Printf("Actual Private Key Type: %T\n", privateKey)
		return nil, fmt.Errorf("invalid private key type, expected *rsa.PrivateKey")
	}
	return &RSASigner{privateKey: rsaPrivateKey}, nil
}

// Sign signs the provided data using RSA.
func (s *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	hashed := sha256.Sum256(dataToBeSigned)
	return rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hashed[:])
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
		eccEncoded, errGenerator := eccGenerator.Generate()
		if errGenerator != nil {
			err = errGenerator
			return
		}
		marshaller := ECCMarshaler{}
		return marshaller.Encode(*eccEncoded)
	default:
		err = fmt.Errorf("wrong algorithm Type %s", algorithm)
		return
	}
}
