package domain

import (
	"sync"
	"time"
)

// SignatureDevice represents a signature device
type SignatureDevice struct {
	ID                 string       `gorm:"primaryKey" json:"id"`
	Label              string       `gorm:"not null" json:"label"`
	SignatureAlgorithm string       `gorm:"not null" json:"signature_alg" validate:"required,oneof='RSA' 'ECC'"`
	PrivateKey         EncryptedKey `gorm:"embedded" json:"-"`
	PublicKey          []byte       `gorm:"not null" json:"public_key"`
	SignatureCounter   SafeCounter  `gorm:"not null" json:"signature_counter"`
	CreatedAt          time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

// EncryptedKey represents an encrypted private key
type EncryptedKey struct {
	Algorithm string `gorm:"column:algorithm" json:"algorithm"`
	Data      []byte `gorm:"column:data" json:"-"`
}

type CreateSignatureDeviceRequest struct {
	Algorithm string `json:"algorithm"`
	Label     string `json:"label"`
}

type SignatureResponse struct {
	Signature   string `json:"signature"`
	SignedData  string `json:"signed_data"`
	DeviceLabel string `json:"device_label"`
}

type CreateSignatureDeviceResponse struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

const (
	RSA string = "RSA"
	ECC string = "ECC"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) IncrementCounter(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}
