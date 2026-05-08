package repository

import (
	"errors"
	"log"

	"github.com/kyeeego/latexmes/domain"
	"gorm.io/gorm"
)

type connectionRepository struct {
	db *gorm.DB
}

func newConnectionRepository(db *gorm.DB) *connectionRepository {
	return &connectionRepository{db}
}

func (r *connectionRepository) GetById(id uint) (*domain.Connection, error) {
	var conn *domain.Connection
	res := r.db.Find(&conn, id)

	return conn, res.Error
}

func (r *connectionRepository) GetUserConnections(id uint) ([]domain.Connection, error) {
	var conns []domain.Connection
	res := r.db.Where("U1 = ? OR U2 = ?", id, id).Find(&conns)

	return conns, res.Error
}

func (r *connectionRepository) CreateConnection(u1 uint, u2 uint) error {
	conn := &domain.Connection{
		U1: u1,
		U2: u2,
	}

	res := r.db.Create(conn)
	return res.Error
}

func (r *connectionRepository) Exists(u1 uint, u2 uint) (bool, error) {
	var conn *domain.Connection
	err := r.db.Where("U1 = ? AND U2 = ? OR U1 = ? AND U2 = ?", u1, u2, u2, u1).First(&conn).Error
	log.Println(err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
