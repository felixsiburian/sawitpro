package repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/model"
	"time"
)

func (r *Repository) Create(ctx context.Context, payload model.Estate) error {
	query := `
		INSERT INTO 
			estate (id, width, length, created_at, updated_at, deleted_at) 
		VALUES 
			($1, $2, $3, $4, $5, $6)`

	stmt, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload.Id, payload.Width, payload.Length, time.Now(), time.Now(), nil)
	if err != nil {
		return err
	}

	return err
}
