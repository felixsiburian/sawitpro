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

func (s Service) CreateTree(ctx context.Context, payload PayloadCreateTree) (uuid.UUID, error) {
	id := uuid.New()
	payload.Id = id

	if err := utils.Validator(payload); err != nil {
		return uuid.Nil, err
	}

	estateData, err := s.repo.FindEstateById(ctx, payload.EstateId)
	if err != nil {
		return uuid.Nil, err
	}

	if payload.Width <= estateData.Width && payload.Length <= estateData.Length {
		if err := s.repo.CreateTree(ctx, model.Tree{
			Id:       id,
			Width:    payload.Width,
			Length:   payload.Length,
			Height:   payload.Height,
			EstateId: payload.EstateId,
		}); err != nil {
			return uuid.Nil, err
		}
	} else {
		return uuid.Nil, err
	}

	return id, nil
}
