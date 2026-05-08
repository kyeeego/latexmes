package repository

import (
	"github.com/kyeeego/latexmes/domain"
	"gorm.io/gorm"
)

type mesreqRepository struct {
	db *gorm.DB
}

func newMesreqRepository(db *gorm.DB) *mesreqRepository {
	return &mesreqRepository{db}
}

func (r *mesreqRepository) GetById(id uint) (*domain.MessageRequest, error) {
	var conn *domain.MessageRequest
	res := r.db.Find(&conn, id)

	return conn, res.Error
}

func (r *mesreqRepository) GetByUser(id uint) ([]domain.MessageRequest, error) {
	var conns []domain.MessageRequest
	res := r.db.Where("to = ?", id).Find(&conns)

	return conns, res.Error
}

func (r *mesreqRepository) Create(from uint, to uint) (*domain.MessageRequest, error) {
	conn := &domain.MessageRequest{
		From: from,
		To:   to,
	}

	res := r.db.Save(conn)
	return conn, res.Error
}

func (r *mesreqRepository) Delete(id uint) error {
	return r.db.Delete(&domain.MessageRequest{}, id).Error
}
