package service

import (
	"fmt"
	"os"

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
	Login(username, password string) (string, error)
	Authorize(jwt string) (int, error)
}

type MessageService interface {
	GetConnectionMessages(connection uint) ([]domain.Message, error)
	GetUserConnections(userid uint) ([]domain.Connection, error)
	SaveMessage(msg dto.ChatUpdate, status domain.MessageStatus) error
}

type Service struct {
	User UserService
	Chat MessageService
	Auth AuthService
}

func New(repository *repository.Repository) *Service {
	m, err := auth.NewManager(os.Getenv("JWT_KEY"))
	if err != nil {
		panic(fmt.Sprintf("couldnt initialize services: %e\n", err))
	}

	return &Service{
		User: newUserService(repository),
		Auth: newAuthService(m, repository),
	}
}
