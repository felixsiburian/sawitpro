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

	PayloadCreateTreeStats struct {
		TreeId   uuid.UUID `json:"tree_id" validate:"required"`
		EstateId uuid.UUID `json:"estate_id" validate:"required"`
		Width    int       `json:"width" validate:"required,min=1,max=50000"`
		Length   int       `json:"length" validate:"required,min=1,max=50000"`
		Height   int       `json:"height" validate:"required,min=1,max=30"`
	}

	PayloadUpdateTree struct {
		Id     uuid.UUID `json:"id" validate:"required"`
		Width  int       `json:"width" validate:"min=1,max=50000"`
		Length int       `json:"length" validate:"min=1,max=50000"`
		Height int       `json:"height" validate:"min=1,max=30"`
	}

	TreeStatsByEstateIdResponse struct {
		Count  int `json:"count"`
		Max    int `json:"max"`
		Min    int `json:"min"`
		Median int `json:"median"`
	}
)
