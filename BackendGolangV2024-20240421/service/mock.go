package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) CreateEstate(ctx context.Context, payload PayloadCreateEstate) (uuid.UUID, error) {
	args := s.Called(ctx, payload)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (s *MockService) CreateTree(ctx context.Context, payload PayloadCreateTree) (uuid.UUID, error) {
	args := s.Called(ctx, payload)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (s *MockService) UpdateTree(ctx context.Context, payload PayloadUpdateTree) (uuid.UUID, error) {
	args := s.Called(ctx, payload)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (s *MockService) TreeStatsByEstateId(ctx context.Context, estateId uuid.UUID) (TreeStatsByEstateIdResponse, error) {
	args := s.Called(ctx, estateId)
	return args.Get(0).(TreeStatsByEstateIdResponse), args.Error(1)
}

func (s *MockService) DroneDistance(ctx context.Context, estateId uuid.UUID) (int, error) {
	args := s.Called(ctx, estateId)
	return args.Get(0).(int), args.Error(1)
}
