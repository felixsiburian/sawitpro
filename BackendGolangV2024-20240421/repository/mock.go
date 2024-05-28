package repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) Create(ctx context.Context, payload model.Estate) error {
	args := r.Called(ctx, payload)
	return args.Error(0)
}

func (r *MockRepo) CreateTree(ctx context.Context, payload model.Tree) error {
	args := r.Called(ctx, payload)
	return args.Error(0)
}

func (r *MockRepo) FindEstateById(ctx context.Context, id uuid.UUID) (model.Estate, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(model.Estate), args.Error(1)
}

func (r *MockRepo) CreateStats(ctx context.Context, payload model.Stats) error {
	args := r.Called(ctx, payload)
	return args.Error(0)
}

func (r *MockRepo) UpdateTree(ctx context.Context, payload model.Tree) error {
	args := r.Called(ctx, payload)
	return args.Error(0)
}

func (r *MockRepo) FindTreeById(ctx context.Context, id uuid.UUID) (model.Tree, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(model.Tree), args.Error(1)
}

func (r *MockRepo) FindStatsByEstateId(ctx context.Context, id uuid.UUID) (FindStatsResponse, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(FindStatsResponse), args.Error(1)
}

func (r *MockRepo) ListStatsByEstateId(ctx context.Context, id uuid.UUID) ([]model.Stats, error) {
	args := r.Called(ctx, id)
	return args.Get(0).([]model.Stats), args.Error(1)
}

func (r *MockRepo) FindAllTreeByEstateId(ctx context.Context, estateId uuid.UUID) ([]model.Tree, error) {
	args := r.Called(ctx, estateId)
	return args.Get(0).([]model.Tree), args.Error(1)
}
