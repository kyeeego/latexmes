package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kyeeego/latexmes/domain"
	"github.com/kyeeego/latexmes/domain/dto"
	"github.com/kyeeego/latexmes/pkg/auth"
	"github.com/kyeeego/latexmes/repository"
	"gorm.io/gorm"
)

type SessionServiceImpl struct {
	repository   *repository.Repository
	tokenManager auth.TokenManager
}

func newSessionService(r *repository.Repository) *SessionServiceImpl {
	return &SessionServiceImpl{
		repository: r,
	}
}

func (s *SessionServiceImpl) Create(user uint, fingerprint string) (string, error) {
	refreshToken := uuid.New().String()

	session := &domain.Session{
		UserID:       user,
		ExpiresAt:    uint64(time.Now().Add(time.Hour * 24 * 30).UnixNano()),
		Fingerprint:  fingerprint,
		RefreshToken: refreshToken,
	}

	err := s.repository.Session.Save(session)

	return refreshToken, err
}

func (s *SessionServiceImpl) Renew(refreshToken string, fingerprint string) (*dto.TokenPair, error) {
	session, err := s.repository.Session.FindByRefreshToken(refreshToken)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Unauthorized")
	} else if err != nil {
		return nil, err
	}

	err = s.repository.Session.Delete(session)
	if err != nil {
		return nil, err
	}

	if session.Fingerprint != fingerprint {
		return nil, errors.New("Invalid fingerprint")
	}

	if time.Now().After(time.Unix(int64(session.ExpiresAt), 0)) {
		return nil, errors.New("Session expired")
	}

	newRefreshToken := uuid.New().String()
	newJwt, err := s.tokenManager.Sign(int(session.UserID), time.Hour*24)
	if err != nil {
		return nil, err
	}

	newSession := &domain.Session{
		UserID:       session.UserID,
		ExpiresAt:    uint64(time.Now().Add(time.Hour * 24 * 30).UnixNano()),
		Fingerprint:  fingerprint,
		RefreshToken: newRefreshToken,
	}

	err = s.repository.Session.Save(newSession)

	return &dto.TokenPair{
		Jwt:          newJwt,
		RefreshToken: newRefreshToken,
	}, err
}

func (s *SessionServiceImpl) Logout(fingerprint string) error {
	return s.repository.Session.DeleteByFingerprint(fingerprint)
}
