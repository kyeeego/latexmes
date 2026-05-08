package service

import (
	"github.com/kyeeego/latexmes/domain"
	"github.com/kyeeego/latexmes/repository"
)

type MsgReqServiceImpl struct {
	repository *repository.Repository
}

func newMsgreqService(r *repository.Repository) *MsgReqServiceImpl {
	return &MsgReqServiceImpl{
		repository: r,
	}
}

func (s *MsgReqServiceImpl) Create(u1 uint, u2 uint) (*domain.MessageRequest, error) {
	return s.repository.MesReq.Create(u1, u2)
}

func (s *MsgReqServiceImpl) GetById(id uint) (*domain.MessageRequest, error) {
	return s.repository.MesReq.GetById(id)
}

func (s *MsgReqServiceImpl) GetByUser(userid uint) ([]domain.MessageRequest, error) {
	return s.repository.MesReq.GetByUser(userid)
}

func (s *MsgReqServiceImpl) Delete(userid uint) error {
	return s.repository.MesReq.Delete(userid)
}
