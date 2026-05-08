package service

import (
	"errors"
	"time"

	"github.com/kyeeego/latexmes/domain/dto"
	"github.com/kyeeego/latexmes/pkg/auth"
	"github.com/kyeeego/latexmes/repository"
)

type AuthServiceImpl struct {
	tokenManager auth.TokenManager
	repository   *repository.Repository
	session      SessionService
}

func newAuthService(manager auth.TokenManager, user *repository.Repository, s SessionService) *AuthServiceImpl {
	return &AuthServiceImpl{
		tokenManager: manager,
		repository:   user,
		session:      s,
	}
}

func (s *AuthServiceImpl) Login(body dto.LoginDto) (*dto.TokenPair, error) {
	user, err := s.repository.User.GetByUsername(body.Username)
	if err != nil {
		return nil, err
	}

	if !auth.Verify(body.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	// refreshToken, err := s.session.Create(user.ID, body.Fingerprint)
	// if err != nil {
	// 	return nil, err
	// }

	jwt, err := s.tokenManager.Sign(int(user.ID), time.Hour*24)
	if err != nil {
		return nil, err
	}

	return &dto.TokenPair{
		Jwt:          jwt,
		RefreshToken: "",
	}, nil
}

func (s *AuthServiceImpl) Authorize(jwt string) (int, error) {
	return s.tokenManager.Verify(jwt)
}
