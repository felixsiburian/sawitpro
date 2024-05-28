package service_test

import (
	"context"
	"github.com/SawitProRecruitment/UserService/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mockService = new(service.MockService)
	estateId    = uuid.MustParse("f841a9a9-5774-4dcc-9b75-6945bab9025a")
	treeId      = uuid.MustParse("e65fdd0c-548c-4a53-9ee6-c08b4b672a4c")
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	newData := service.PayloadCreateEstate{
		Id:     estateId,
		Width:  5,
		Length: 10,
	}

	mockService.On("CreateEstate", ctx, newData).Return(estateId, nil)

	result, err := mockService.CreateEstate(ctx, newData)

	mockService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, estateId, result)
}

func TestCreateTree(t *testing.T) {
	ctx := context.Background()
	newData := service.PayloadCreateTree{
		Id:       treeId,
		EstateId: estateId,
		Width:    3,
		Length:   4,
		Height:   3,
	}

	mockService.On("CreateTree", ctx, newData).Return(treeId, nil)

	result, err := mockService.CreateTree(ctx, newData)
	mockService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, treeId, result)
}

func TestUpdateTree(t *testing.T) {
	ctx := context.Background()
	updateData := service.PayloadUpdateTree{
		Id:     treeId,
		Width:  4,
		Length: 4,
		Height: 4,
	}

	mockService.On("UpdateTree", ctx, updateData).Return(treeId, nil)

	result, err := mockService.UpdateTree(ctx, updateData)
	mockService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, treeId, result)
}

func TestTreeStatsByEstateId(t *testing.T) {
	ctx := context.Background()
	expectedTree := service.TreeStatsByEstateIdResponse{
		Count:  1,
		Max:    4,
		Min:    3,
		Median: 3,
	}

	mockService.On("TreeStatsByEstateId", ctx, estateId).Return(expectedTree, nil)

	result, err := mockService.TreeStatsByEstateId(ctx, estateId)

	mockService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedTree, result)
}

func TestDroneDistance(t *testing.T) {
	ctx := context.Background()
	expected := 4

	mockService.On("DroneDistance", ctx, estateId).Return(expected, nil)

	result, err := mockService.DroneDistance(ctx, estateId)

	mockService.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
