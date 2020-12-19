package store_test

import (
	"github.com/GlebSolncev/http-rest-api/internal/app/models"
	"github.com/GlebSolncev/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationReposiroty_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("applications")

	u, err := s.Application().Create(&models.Application{
		Body:   `{"status": "work"}`,
		Status: "in_process",
		Slug:   "Demo",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// TestApplicationRepository_Find ...
func TestApplicationRepository_Find(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("applications")
	id := 999
	_, err := s.Application().Find(id)
	assert.Error(t, err)

	newApp, errorApp := s.Application().Create(&models.Application{
		//Id: 1,
		Body:   `{"status": "work"}`,
		Status: "in_process",
		Slug:   "Demo",
	})
	assert.NoError(t, errorApp)
	assert.NotNil(t, newApp)

	u, err := s.Application().Find(newApp.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// TestApplicationRepository_Find ...
func TestApplicationRepository_FindBy(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("applications")
	field := "slug"
	slug := "hello-world"
	_, err := s.Application().FindBy(field, slug)
	assert.Error(t, err)

	newApp, errorApp := s.Application().Create(&models.Application{
		Body:   `{"status": "work"}`,
		Status: "in_process",
		Slug:   "demo",
	})
	assert.NoError(t, errorApp)
	assert.NotNil(t, newApp.Id)

	u, err := s.Application().FindBy("slug", newApp.Slug)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	ByUknownName, err := s.Application().FindBy("FieldUnknown", newApp.Slug)
	assert.Error(t, err)
	assert.Nil(t, ByUknownName)
}

// TestApplicationRepository_Update ...
func TestApplicationRepository_Update(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("applications")

	a, errorApp := s.Application().Create(&models.Application{
		Body:   `{"status": "work"}`,
		Status: "in_process",
		Slug:   "demo",
	})
	assert.NoError(t, errorApp)
	assert.NotNil(t, a.Id)

	a.Slug = "demo-v2"
	a.Body = `{"version": 2}`
	a.Status = "done"

	_, _ = s.Application().Update(a)
}
