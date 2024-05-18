package service

import (
	"context"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	CreateEstate(ctx context.Context, payload PayloadCreateEstate) (uuid.UUID, error)
	CreateTree(ctx context.Context, payload PayloadCreateTree) (uuid.UUID, error)
	UpdateTree(ctx context.Context, payload PayloadUpdateTree) (uuid.UUID, error)
	TreeStatsByEstateId(ctx context.Context, estateId uuid.UUID) (TreeStatsByEstateIdResponse, error)
}
