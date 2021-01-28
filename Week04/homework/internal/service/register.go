package service

import (
	"github.com/geekbosszhang/Go-001/Week04/internal/data"
)

type registerService struct {
	user data.User
}

func NewRegisterService() *registerService {
	return &registerService{}
}

func (s *registerService) Register(user data.User) {
	return user
}