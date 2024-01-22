package service

import (
	"test_task/internal/models"
	"test_task/internal/repository"
)

type People interface {
	CreatePerson(newUser models.Person) (int, error)
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
