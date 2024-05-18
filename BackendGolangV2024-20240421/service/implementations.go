package service

import (
	"context"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/google/uuid"
)

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateEstate(ctx context.Context, payload PayloadCreateEstate) (uuid.UUID, error) {
	id := uuid.New()
	payload.Id = id

	if err := utils.Validator(payload); err != nil {
		return uuid.Nil, err
	}

	if err := s.repo.Create(ctx, model.Estate{
		Id:     payload.Id,
		Width:  payload.Width,
		Length: payload.Length,
	}); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
