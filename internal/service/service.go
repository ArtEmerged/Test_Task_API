package service

import (
	"test_task/internal/models"
	"test_task/internal/repository"
)

type People interface {
	CreatePerson(newPerson models.Person) (int, error)
	GetPersonById(id int) (models.Person, error)
	GetPeople(filters models.Filters) ([]models.Person, error)
	UpdatePerson(id int, person models.Person) error
	DeletePerson(id int) error
}

type Service struct {
	People
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		People: NewPeopleService(repo),
	}
}
