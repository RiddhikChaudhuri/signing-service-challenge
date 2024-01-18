package domain

import "time"

// SignatureDevice represents a signature device
type Signature struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	DeviceID        string    `gorm:"not null" json:"deviceId"`
	SignatureBase64 string    `gorm:"not null" json:"signature"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type SignatureRequest struct {
	DeviceID string `json:"deviceId"`
	Data     string `json:"data"`
}
