package service

import (
	"test_task/internal/models"
	"test_task/internal/repository"
)

type People interface {
	CreatePerson(newPerson models.Person) (int, error)
	DeletePerson(id int) error
	UpdatePerson(id int, person models.Person) error
	GetPersonById(id int) (models.Person, error)
}

type Service struct {
	People
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		People: NewPeopleService(repo),
	}
}
