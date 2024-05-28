package repository_test

import (
	"context"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	mockRepo = new(repository.MockRepo)
	estateId = uuid.MustParse("f841a9a9-5774-4dcc-9b75-6945bab9025a")
	treeId   = uuid.MustParse("e65fdd0c-548c-4a53-9ee6-c08b4b672a4c")
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	newData := model.Estate{
		Id:        estateId,
		Width:     10,
		Length:    20,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	mockRepo.On("Create", ctx, newData).Return(nil)

	err := mockRepo.Create(ctx, newData)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
}

func TestCreateTree(t *testing.T) {
	ctx := context.Background()
	newData := model.Tree{
		Id:        treeId,
		Width:     3,
		Length:    4,
		Height:    3,
		EstateId:  estateId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	mockRepo.On("CreateTree", ctx, newData).Return(nil)

	err := mockRepo.CreateTree(ctx, newData)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
}

func TestFindEstateById(t *testing.T) {
	ctx := context.Background()
	expected := model.Estate{
		Id:     estateId,
		Width:  10,
		Length: 20,
	}

	mockRepo.On("FindEstateById", ctx, estateId).Return(expected, nil)
	res, err := mockRepo.FindEstateById(ctx, estateId)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestCreateStats(t *testing.T) {
	ctx := context.Background()
	newData := model.Stats{
		TreeId:    treeId,
		Width:     4,
		Length:    4,
		Height:    4,
		EstateId:  estateId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("CreateStats", ctx, newData).Return(nil)

	err := mockRepo.CreateStats(ctx, newData)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
}

func TestUpdateTree(t *testing.T) {
	ctx := context.Background()
	updatedData := model.Tree{
		Id:        treeId,
		Width:     4,
		Length:    4,
		Height:    6,
		EstateId:  estateId,
		UpdatedAt: time.Now(),
	}

	mockRepo.On("UpdateTree", ctx, updatedData).Return(nil)
	err := mockRepo.UpdateTree(ctx, updatedData)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestFindTreeById(t *testing.T) {
	ctx := context.Background()

	expectedTree := model.Tree{
		Id:       treeId,
		Width:    4,
		Length:   4,
		Height:   6,
		EstateId: estateId,
	}

	mockRepo.On("FindTreeById", ctx, treeId).Return(expectedTree, nil)

	res, err := mockRepo.FindTreeById(ctx, treeId)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedTree, res)
}

func TestFindStatsByEstateId(t *testing.T) {
	ctx := context.Background()

	expected := repository.FindStatsResponse{
		Count: 1,
		Max:   6,
		Min:   4,
	}

	mockRepo.On("FindStatsByEstateId", ctx, estateId).Return(expected, nil)

	res, err := mockRepo.FindStatsByEstateId(ctx, estateId)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestListStatsByEstateId(t *testing.T) {
	ctx := context.Background()

	expected := []model.Stats{
		{
			TreeId:   treeId,
			Width:    4,
			Length:   4,
			Height:   6,
			EstateId: estateId,
		},
	}

	mockRepo.On("ListStatsByEstateId", ctx, estateId).Return(expected, nil)

	res, err := mockRepo.ListStatsByEstateId(ctx, estateId)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestFindAllTreeByEstateId(t *testing.T) {
	ctx := context.Background()

	expected := []model.Tree{
		{
			Id:       treeId,
			Width:    4,
			Length:   4,
			Height:   6,
			EstateId: estateId,
		},
	}

	mockRepo.On("FindAllTreeByEstateId", ctx, estateId).Return(expected, nil)

	res, err := mockRepo.FindAllTreeByEstateId(ctx, estateId)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
