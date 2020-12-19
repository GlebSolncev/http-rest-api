package models

const ApplicationTable = "applications"

type Application struct {
	Id     int
	Slug   string
	Body   string `json:"body"`
	Status string
}
