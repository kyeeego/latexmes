package repository

import (
	"strings"

	"github.com/kyeeego/latexmes/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r userRepository) Insert(model *domain.User) error {
	res := r.db.Create(model)
	return res.Error
}

func (r userRepository) GetById(id uint) (domain.User, error) {
	user := domain.User{}
	res := r.db.First(&user, id)

	return user, res.Error
}

func (r userRepository) GetByUsername(username string) (domain.User, error) {
	user := domain.User{}
	res := r.db.Where("username = ?", username).First(&user)

	return user, res.Error
}

func (r userRepository) Search(q string) ([]domain.User, error) {
	users := []domain.User{}
	res := r.db.Where("LOWER(username) LIKE ?", "%"+strings.ToLower(q)+"%").Find(&users)

	return users, res.Error
}
