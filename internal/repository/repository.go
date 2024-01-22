package repository

import (
	"database/sql"
	"test_task/internal/models"
)

const (
	peopleTable      = "people"
	nationalizeTable = "nationalize"
)

type People interface {
	CreatePerson(newPerson models.Person) (int, error)
	GetPerson(newPerson models.Person) (models.Person, error)
}

type Repository struct {
	People
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{People: NewPeoplePostgres(db)}
}