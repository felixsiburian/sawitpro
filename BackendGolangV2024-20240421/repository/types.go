// This file contains types that are used in the repository layer.
package repository

import "github.com/google/uuid"

type Estate struct {
	Id uuid.UUID `json:"id"`
}

type GetTestByIdOutput struct {
	Name string
}
