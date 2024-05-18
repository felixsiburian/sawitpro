package repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
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

func (r *Repository) CreateTree(ctx context.Context, payload model.Tree) error {

	query := `
		INSERT INTO 
		  	tree (id, width, length, height, estate_id, created_at, updated_at, deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	stmt, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload.Id, payload.Width, payload.Length, payload.Height, payload.EstateId, time.Now(), time.Now(), nil)
	if err != nil {
		return err
	}

	return err
}

func (r *Repository) FindEstateById(ctx context.Context, id uuid.UUID) (model.Estate, error) {
	var res model.Estate

	query := `
		SELECT 
			id, width, length
		FROM
			estate
		where id = $1
	`

	row := r.Db.QueryRow(query, id)
	if err := row.Scan(&res.Id, &res.Width, &res.Length); err != nil {
		return res, err
	}

	return res, nil
}
