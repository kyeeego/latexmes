package service

import (
	"errors"
	"time"

	"github.com/kyeeego/latexmes/pkg/auth"
	"github.com/kyeeego/latexmes/repository"
)

type AuthServiceImpl struct {
	tokenManager auth.TokenManager
	repository   *repository.Repository
}

func newAuthService(manager auth.TokenManager, user *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{
		tokenManager: manager,
		repository:   user,
	}
}

func (s *AuthServiceImpl) Login(username, pass string) (string, error) {
	user, err := s.repository.User.GetByUsername(username)
	if err != nil {
		return "", err
	}

	if !auth.Verify(pass, user.Password) {
		return "", errors.New("invalid password")
	}

	return s.tokenManager.Sign(int(user.ID), time.Hour*24)
}

func (s *AuthServiceImpl) Authorize(jwt string) (int, error) {
	return s.tokenManager.Verify(jwt)
}
