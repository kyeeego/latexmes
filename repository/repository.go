package repository

import (
	"github.com/kyeeego/latexmes/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Insert(model *domain.User) error
}

type ConnectionRepository interface {
	GetUserConnections(id uint) ([]domain.Connection, error)
	CreateConnection(u1 uint, u2 uint) error
	Exists(u1 uint, u2 uint) (bool, error)
}

type MessageRepository interface {
	SaveMessage(msg domain.Message) error
	GetConnectionMessages(connection uint) ([]domain.Message, error)
	UpdateStatus(id uint, status domain.MessageStatus) error
}

type Repository struct {
	User       UserRepository
	Message    MessageRepository
	Connection ConnectionRepository
}

func New(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Connection{})
	db.AutoMigrate(&domain.Message{})

	return &Repository{
		Connection: newConnectionRepository(db),
		Message:    newMessageRepository(db),
		User:       newUserRepository(db),
	}
}
