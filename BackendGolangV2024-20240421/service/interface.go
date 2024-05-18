package service

import (
	"context"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	CreateEstate(ctx context.Context, payload PayloadCreateEstate) (uuid.UUID, error)
}
