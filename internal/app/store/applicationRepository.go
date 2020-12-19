package store

import (
	"fmt"
	"github.com/GlebSolncev/http-rest-api/internal/app/models"
)

type ApplicationRepository struct {
	store *Store
}

// Create ...
func (r *ApplicationRepository) Create(a *models.Application) (*models.Application, error) {
	queryWithTable := fmt.Sprintf("insert into %s (id, slug, body, status) values(null, ?, ?, ?)", models.ApplicationTable)
	stmt, err := r.store.db.Prepare(queryWithTable)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(a.Slug, a.Body, a.Status)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	a.Id = int(id)
	return a, nil
}

// Find ...
func (r *ApplicationRepository) Find(id int) (*models.Application, error) {
	a := &models.Application{}
	query := fmt.Sprintf("select * from %s where id=%d", models.ApplicationTable, id)
	if err := r.store.db.QueryRow(query).Scan(
		&a.Id,
		&a.Slug,
		&a.Body,
		&a.Status,
	); err != nil {
		return nil, err
	}

	return a, nil
}

// FindBy ...
func (r *ApplicationRepository) FindBy(field string, value string) (*models.Application, error) {
	a := &models.Application{}
	query := fmt.Sprintf(`select * from %s where %s = "%s"`, models.ApplicationTable, field, value)

	if err := r.store.db.QueryRow(query).Scan(
		&a.Id,
		&a.Slug,
		&a.Body,
		&a.Status,
	); err != nil {
		return nil, err
	}

	return a, nil
}

// Update ...
func (r *ApplicationRepository) Update(a *models.Application) (*models.Application, error) {
	//u := &models.Application{}
	query := fmt.Sprintf(`update %s set slug="%s", body='%s', status="%s" where id=%d`,
		models.ApplicationTable,
		a.Slug,
		a.Body,
		a.Status,
		a.Id,
	)

	if err := r.store.db.QueryRow(query).Scan(
		&a.Id,
		&a.Slug,
		&a.Body,
		&a.Status,
	); err != nil {
		return nil, err
	}

	return a, nil
}
