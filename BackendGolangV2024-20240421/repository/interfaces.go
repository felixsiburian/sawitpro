// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

type RepositoryInterface interface {
	Create(ctx context.Context, payload model.Estate) error
	CreateTree(ctx context.Context, payload model.Tree) error
	FindEstateById(ctx context.Context, id uuid.UUID) (model.Estate, error)
}
