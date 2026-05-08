package repository

import (
	"github.com/kyeeego/latexmes/domain"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func newNotificationRepository(db *gorm.DB) *notificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) GetUnsent(id uint) ([]domain.Notification, error) {
	var conns []domain.Notification
	res := r.db.Where("recipient = ?", id).Find(&conns)

	return conns, res.Error
}

func (r *notificationRepository) Save(not *domain.Notification) error {
	res := r.db.Save(not)
	return res.Error
}

func (r *notificationRepository) Delete(not *domain.Notification) error {
	return r.db.Delete(not).Error
}
