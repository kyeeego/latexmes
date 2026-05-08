package service

import (
	"github.com/kyeeego/latexmes/domain"
	"github.com/kyeeego/latexmes/domain/dto"
	"github.com/kyeeego/latexmes/repository"
)

type ChatServiceImpl struct {
	repository *repository.Repository
}

func newChatService(r *repository.Repository) *ChatServiceImpl {
	return &ChatServiceImpl{
		repository: r,
	}
}

func (s *ChatServiceImpl) GetConnectionMessages(connection uint) ([]domain.Message, error) {
	return s.repository.Message.GetConnectionMessages(connection)
}

func (s *ChatServiceImpl) SaveMessage(msg dto.ChatUpdate, sender uint) (*domain.Message, error) {
	model := &domain.Message{
		Status:       domain.StatusUnsent,
		Content:      msg.Content,
		ConnectionID: msg.Connection,
		SenderID:     sender,
	}

	err := s.repository.Message.SaveMessage(model)

	return model, err
}
