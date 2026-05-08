package service

import (
	"github.com/kyeeego/latexmes/domain"
	"github.com/kyeeego/latexmes/repository"
)

type ConnectionServiceImpl struct {
	repository *repository.Repository
}

func newConnectionService(r *repository.Repository) *ConnectionServiceImpl {
	return &ConnectionServiceImpl{
		repository: r,
	}
}

func (s *ConnectionServiceImpl) Create(u1 uint, u2 uint) error {
	return s.repository.Connection.CreateConnection(u1, u2)
}

func (s *ConnectionServiceImpl) Exists(u1 uint, u2 uint) (bool, error) {
	return s.repository.Connection.Exists(u1, u2)
}

func (s *ConnectionServiceImpl) GetById(id uint) (*domain.Connection, error) {
	return s.repository.Connection.GetById(id)
}

func (s *ConnectionServiceImpl) GetUserConnections(userid uint) ([]domain.Connection, error) {
	return s.repository.Connection.GetUserConnections(userid)
}
