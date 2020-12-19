package store

import (
	"github.com/GlebSolncev/http-rest-api/internal/app/models"
)

type ApplicationRepository struct {
	store *Store
}

// Create ...
func (r *ApplicationRepository) Create(u *models.Application) (*models.Application, error) {
	stmt, err := r.store.db.Prepare("insert into applications(id, slug, body, status) values(null, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(u.Slug, u.Body, u.Status)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.Id = int(id)
	return u, nil
}

// Find ...
func (r *ApplicationRepository) Find(id int) (*models.Application, error) {
	u := &models.Application{}
	if err := r.store.db.QueryRow(
		"select * from applications where id=?", id,
	).Scan(
		&u.Id,
		&u.Slug,
		&u.Body,
		&u.Status,
	); err != nil {
		return nil, err
	}

	return u, nil
}
