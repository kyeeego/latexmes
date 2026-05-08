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

// TODO thgis is disgusting
func (r *messageRepository) SaveMessage(msg *domain.Message) error {
	r.db.Model(&domain.Message{}).Save(msg)
	var mdl *domain.Message
	res := r.db.Model(&domain.Message{}).Preload("Connection").Find(&mdl, msg.ID)
	msg.Connection = mdl.Connection
	return res.Error
}

func (r *messageRepository) GetConnectionMessages(connection uint) ([]domain.Message, error) {
	var messages []domain.Message
	res := r.db.Where("connection_id = ?", connection).Preload("Connection").Find(&messages)

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
