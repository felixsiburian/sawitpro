package service

import "github.com/google/uuid"

type (
	PayloadCreateEstate struct {
		Id     uuid.UUID `json:"id" validate:"required"`
		Width  int       `json:"width" validate:"required,min=1,max=50000"`
		Length int       `json:"length" validate:"required,min=1,max=50000"`
	}

	PayloadCreateTree struct {
		Id       uuid.UUID `json:"id" validate:"required"`
		EstateId uuid.UUID `json:"estate_id" validate:"required"`
		Width    int       `json:"width" validate:"required,min=1,max=50000"`
		Length   int       `json:"length" validate:"required,min=1,max=50000"`
		Height   int       `json:"height" validate:"required,min=1,max=30"`
	}
)
