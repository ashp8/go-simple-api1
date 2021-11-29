package service

import (
	"context"

	"example.com/ashp8/memrizr/model"
	"github.com/google/uuid"
)

type UserService struct {
	UserRespository model.UserRepository
}

type USConfig struct {
	UserRespository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRespository: c.UserRespository,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRespository.FindById(ctx, uid)

	return u, err
}
