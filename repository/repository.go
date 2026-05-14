package repository

import (
	"github.com/kyeeego/latexmes/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Insert(model *domain.User) error
	Search(q string) ([]domain.User, error)
}

type ConnectionRepository interface {
	GetById(id uint) (*domain.Connection, error)
	GetUserConnections(id uint) ([]domain.Connection, error)
	CreateConnection(u1 uint, u2 uint) error
	Exists(u1 uint, u2 uint) (bool, error)
}

type MessageRepository interface {
	SaveMessage(msg *domain.Message) error
	GetConnectionMessages(connection uint) ([]domain.Message, error)
	UpdateStatus(id uint, status domain.MessageStatus) error
}

type MessageRequestRepository interface {
	GetById(id uint) (*domain.MessageRequest, error)
	GetByUser(id uint) ([]domain.MessageRequest, error)
	Create(from uint, to uint) (*domain.MessageRequest, error)
	Delete(id uint) error
}

type NotificationRepository interface {
	Save(*domain.Notification) error
	GetUnsent(userid uint) ([]domain.Notification, error)
	Delete(*domain.Notification) error
}

type SessionRepository interface {
	Save(*domain.Session) error
	Delete(*domain.Session) error
	FindByUser(id uint) ([]*domain.Session, error)
	FindByRefreshToken(token string) (*domain.Session, error)
	DeleteByUser(id uint) error
	DeleteByFingerprint(fingerprint string) error
}

type Repository struct {
	User         UserRepository
	Message      MessageRepository
	Connection   ConnectionRepository
	MesReq       MessageRequestRepository
	Session      SessionRepository
	Notification NotificationRepository
}

func New(db *gorm.DB) *Repository {
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Connection{})
	db.AutoMigrate(&domain.Message{})
	db.AutoMigrate(&domain.MessageRequest{})
	db.AutoMigrate(&domain.Notification{})
	// db.AutoMigrate(&domain.Session{})

	return &Repository{
		Connection:   newConnectionRepository(db),
		Message:      newMessageRepository(db),
		User:         newUserRepository(db),
		MesReq:       newMesreqRepository(db),
		Notification: newNotificationRepository(db),
	}
}
