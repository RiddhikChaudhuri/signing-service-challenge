package domain

import "crypto/rsa"

// SignatureDevice represents a signature device
type SignatureDevice struct {
	DeviceID           string
	Label              string
	SignatureAlgorithm string
	PrivateKey         *rsa.PrivateKey
	PublicKey          []byte
	SignatureCounter   int
}
