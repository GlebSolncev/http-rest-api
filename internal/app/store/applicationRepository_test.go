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
		//Id: 1,
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

	newUser, errorUser := s.Application().Create(&models.Application{
		//Id: 1,
		Body:   `{"status": "work"}`,
		Status: "in_process",
		Slug:   "Demo",
	})
	assert.NoError(t, errorUser)
	assert.NotNil(t, newUser)

	u, err := s.Application().Find(newUser.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
