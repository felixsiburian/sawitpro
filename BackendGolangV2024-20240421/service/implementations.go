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

		if err := s.repo.CreateStats(ctx, model.Stats{
			TreeId:   id,
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

func (s Service) UpdateTree(ctx context.Context, payload PayloadUpdateTree) (uuid.UUID, error) {
	treeData, err := s.repo.FindTreeById(ctx, payload.Id)
	if err != nil {
		return uuid.Nil, err
	}

	if payload.Width != treeData.Width && payload.Width > 0 {
		treeData.Width = payload.Width
	}

	if payload.Length != treeData.Length && payload.Length > 0 {
		treeData.Length = payload.Length
	}

	if payload.Height != treeData.Height && payload.Height > 0 {
		treeData.Height = payload.Height
	}

	if err := s.repo.UpdateTree(ctx, treeData); err != nil {
		return uuid.Nil, err
	}

	if err := s.repo.CreateStats(ctx, model.Stats{
		TreeId:   treeData.Id,
		Width:    treeData.Width,
		Length:   treeData.Length,
		Height:   treeData.Height,
		EstateId: treeData.EstateId,
	}); err != nil {
		return uuid.Nil, err
	}

	return treeData.Id, nil
}

func (s Service) TreeStatsByEstateId(ctx context.Context, estateId uuid.UUID) (TreeStatsByEstateIdResponse, error) {
	var res TreeStatsByEstateIdResponse

	treeStats, err := s.repo.FindStatsByEstateId(ctx, estateId)
	if err != nil {
		return res, err
	}

	res.Count = treeStats.Count
	res.Max = treeStats.Max
	res.Min = treeStats.Min

	median, err := s.countMedianHeight(ctx, estateId)
	if err != nil {
		return res, err
	}

	res.Median = median

	return res, nil
}

func (s Service) countMedianHeight(ctx context.Context, id uuid.UUID) (int, error) {
	var (
		err     error
		heights []int
	)
	statsData, err := s.repo.ListStatsByEstateId(ctx, id)
	if err != nil {
		return -1, err
	}

	for _, v := range statsData {
		heights = append(heights, v.Height)
	}

	x := len(heights)
	if x == 0 {
		return 0, nil
	}

	if x%2 == 1 {
		return heights[x/2], nil
	}

	mid := x / 2
	return (heights[mid-1] + heights[mid]) / 2, nil
}
