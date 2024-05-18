package model

import (
	"github.com/google/uuid"
	"time"
)

// this struct represent db
type (
	Estate struct {
		Id        uuid.UUID  `json:"id"`
		Width     int        `json:"width"`
		Length    int        `json:"length"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	Tree struct {
		Id        uuid.UUID  `json:"id"`
		Width     int        `json:"width"`
		Length    int        `json:"length"`
		Height    int        `json:"height"`
		EstateId  uuid.UUID  `json:"estate_id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	Stats struct {
		TreeId    uuid.UUID  `json:"tree_id"`
		Width     int        `json:"width"`
		Length    int        `json:"length"`
		Height    int        `json:"height"`
		EstateId  uuid.UUID  `json:"estate_id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)
