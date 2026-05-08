package domain

import "gorm.io/gorm"

type NotificationType int

const (
	TypeMessage NotificationType = iota
	TypeMsgRequest
)

// type Notification[T any] struct {
// 	gorm.Model
// 	Type      NotificationType `json:"type"`
// 	Recipient uint             `json:"recipient"`
// 	Content   T                `json:"content" gorm:"serializer:json"`
// }

type Notification struct {
	gorm.Model
	Type      NotificationType `json:"type"`
	Recipient uint             `json:"recipient"`
	Content   any              `json:"content" gorm:"serializer:json"`
}
