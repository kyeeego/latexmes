package service

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kyeeego/latexmes/domain"
	"github.com/kyeeego/latexmes/domain/dto"
	"github.com/kyeeego/latexmes/pkg/auth"
	"github.com/kyeeego/latexmes/repository"
)

type UserService interface {
	Register(username string, password string) error
	GetById(id uint) (domain.UserDto, error)
}

type AuthService interface {
	Login(body dto.LoginDto) (*dto.TokenPair, error)
	Authorize(jwt string) (int, error)
}

type MessageService interface {
	GetConnectionMessages(connection uint) ([]domain.Message, error)
	SaveMessage(msg dto.ChatUpdate, sender uint) (*domain.Message, error)
}

type ConnectionService interface {
	Create(u1 uint, u2 uint) error
	Exists(u1 uint, u2 uint) (bool, error)
	GetById(id uint) (*domain.Connection, error)
	GetUserConnections(userid uint) ([]domain.Connection, error)
}

type MsgReqService interface {
	Create(from uint, to uint) (*domain.MessageRequest, error)
	GetByUser(userid uint) ([]domain.MessageRequest, error)
	GetById(id uint) (*domain.MessageRequest, error)
	Delete(userid uint) error
}

type SessionService interface {
	Create(user uint, fingerprint string) (string, error)
	Renew(refreshToken string, fingerprint string) (*dto.TokenPair, error)
	Logout(fingerprint string) error
}

type S3Service interface {
	UploadFile(uploader uint, filename string, file multipart.File) (string, error)
}

type NotificationService interface {
	Save(*domain.Notification) error
	GetUnsent(userid uint) ([]domain.Notification, error)
	Delete(*domain.Notification) error
}

type Service struct {
	User         UserService
	Chat         MessageService
	Connection   ConnectionService
	Auth         AuthService
	MsgReq       MsgReqService
	Session      SessionService
	S3           S3Service
	Notification NotificationService
}

func New(repository *repository.Repository, s3client *s3.Client) *Service {
	m, err := auth.NewManager(os.Getenv("JWT_KEY"))
	if err != nil {
		panic(fmt.Sprintf("couldnt initialize services: %e\n", err))
	}

	s := newSessionService(repository)

	return &Service{
		User:         newUserService(repository),
		Auth:         newAuthService(m, repository, s),
		Chat:         newChatService(repository),
		Connection:   newConnectionService(repository),
		S3:           newS3Service(s3client, "TODO"),
		MsgReq:       newMsgreqService(repository),
		Notification: newNotificationService(repository),
	}
}
