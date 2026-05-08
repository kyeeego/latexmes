package domain

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	RefreshToken string `json:"refresh_token"`
	UserID       uint   `json:"user_id"`
	ExpiresAt    uint64 `json:"expires_at"`
	Fingerprint  string `json:"fingerprint"`
}
