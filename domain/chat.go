package domain

import "gorm.io/gorm"

type MessageStatus int

const (
	StatusUnsent MessageStatus = iota
	StatusSent
)

type Connection struct {
	gorm.Model
	U1 uint
	U2 uint
}

type Message struct {
	gorm.Model
	Status       MessageStatus `json:"status"`
	Content      string        `json:"content"`
	ConnectionID uint          `json:"connection_id"`
	SenderID     uint          `json:"sender_id"`
	Connection   Connection    `json:"-"`
	Sender       User          `json:"-"`
}
