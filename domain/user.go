package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	// TODO store pfps in yandex cloud bucket
	// Pfp      string `json:"pfp"`
	Password string `json:"password"`
}

type UserDto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	// Pfp string `json:"pfp"`
}
