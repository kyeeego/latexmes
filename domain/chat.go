package domain

import "gorm.io/gorm"

type MessageStatus int

const (
	StatusUnsent MessageStatus = iota
	StatusSent
)

type Connection struct {
	ID uint
	U1 uint
	U2 uint
}

type Message struct {
	gorm.Model
	Status       MessageStatus `json:"status"`
	Content      string        `json:"content"`
	ConnectionID uint          `json:"connection_id"`
	Connection   Connection
	SenderID     uint `json:"sender_id"`
	Sender       User
}

type MessageRequest struct {
	gorm.Model
	From     uint    `json:"from"`
	FromUser UserDto `json:"from_user" gorm:"-"`
	To       uint    `json:"to"`
}
