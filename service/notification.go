package service

import (
	"github.com/kyeeego/latexmes/domain"
	"github.com/kyeeego/latexmes/repository"
)

type NotificationServiceImpl struct {
	repository *repository.Repository
}

func newNotificationService(r *repository.Repository) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		repository: r,
	}
}

func (s *NotificationServiceImpl) Save(not *domain.Notification) error {
	return s.repository.Notification.Save(not)
}

func (s *NotificationServiceImpl) GetUnsent(userid uint) ([]domain.Notification, error) {
	return s.repository.Notification.GetUnsent(userid)
}

func (s *NotificationServiceImpl) Delete(not *domain.Notification) error {
	return s.repository.Notification.Delete(not)
}
