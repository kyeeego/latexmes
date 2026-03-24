package repository

import (
	"github.com/kyeeego/latexmes/domain"
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

func newMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) SaveMessage(msg domain.Message) error {
	res := r.db.Create(msg)
	return res.Error
}

func (r *messageRepository) GetConnectionMessages(connection uint) ([]domain.Message, error) {
	var messages []domain.Message
	res := r.db.Where("connection_id = ?", connection).Find(&messages)

	return messages, res.Error
}

func (r *messageRepository) UpdateStatus(id uint, status domain.MessageStatus) error {
	var msg domain.Message
	err := r.db.First(&msg, id).Error
	if err != nil {
		return err
	}

	msg.Status = status

	return r.db.Save(&msg).Error
}
