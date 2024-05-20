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
		where 
		    id = $1
			AND deleted_at is null
	`

	row := r.Db.QueryRow(query, id)
	if err := row.Scan(&res.Id, &res.Width, &res.Length); err != nil {
		return res, err
	}

	return res, nil
}

func (r *Repository) CreateStats(ctx context.Context, payload model.Stats) error {
	query := ` 
		INSERT INTO 
		  	stats (tree_id, width, length, height, estate_id, created_at, updated_at, deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	stmt, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload.TreeId, payload.Width, payload.Length, payload.Height, payload.EstateId, time.Now(), time.Now(), nil)
	if err != nil {
		return err
	}

	return err
}

func (r *Repository) UpdateTree(ctx context.Context, payload model.Tree) error {
	query := `
		UPDATE
			tree
		SET
			length = $1,
			width = $2,
			height = $3,
			updated_at = $4
		WHERE
			id = $5
	`

	stmt, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload.Length, payload.Width, payload.Height, time.Now(), payload.Id)
	if err != nil {
		return err
	}

	return err
}

func (r *Repository) FindTreeById(ctx context.Context, id uuid.UUID) (model.Tree, error) {
	var res model.Tree

	query := `
		SELECT 
			id, width, length, height, estate_id
		FROM
			tree
		where 
		    id = $1
			AND deleted_at is null
	`

	row := r.Db.QueryRow(query, id)
	if err := row.Scan(&res.Id, &res.Width, &res.Length, &res.Height, &res.EstateId); err != nil {
		return res, err
	}

	return res, nil
}

func (r *Repository) FindStatsByEstateId(ctx context.Context, id uuid.UUID) (FindStatsResponse, error) {
	var res FindStatsResponse

	query := `
		select 
			count(distinct tree_id),max(height), min(height) 
		from 
			stats 
		where 
			estate_id = $1
			AND deleted_at is null
	`

	row := r.Db.QueryRow(query, id)
	if err := row.Scan(&res.Count, &res.Max, &res.Min); err != nil {
		return res, err
	}

	return res, nil
}

func (r *Repository) ListStatsByEstateId(ctx context.Context, id uuid.UUID) ([]model.Stats, error) {
	var res []model.Stats

	query := `
		SELECT
			tree_id, width, length, height
		FROM
		 	stats
		WHERE
			estate_id = $1
			AND deleted_at is null
	`

	rows, err := r.Db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var stats model.Stats
		if err := rows.Scan(&stats.TreeId, &stats.Width, &stats.Length, &stats.Height); err != nil {
			return nil, err
		}

		res = append(res, stats)
	}

	return res, nil
}

func (r *Repository) FindAllTreeByEstateId(ctx context.Context, estateId uuid.UUID) ([]model.Tree, error) {
	var res []model.Tree

	query := `
		select 
			id, width, length, height
		from
			tree
		where
			estate_id = $1
			AND deleted_at is null
	`

	rows, err := r.Db.Query(query, estateId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tree model.Tree
		if err := rows.Scan(&tree.Id, &tree.Width, &tree.Length, &tree.Height); err != nil {
			return nil, err
		}

		res = append(res, tree)
	}

	return res, nil
}
