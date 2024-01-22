package service

import (
	"fmt"
	"test_task/internal/models"
	"test_task/internal/repository"
	"test_task/pkg"
)

type PeopleService struct {
	repo repository.People
}

func NewPeopleService(repo *repository.Repository) *PeopleService {
	return &PeopleService{repo: repo.People}
}

func (s *PeopleService) CreatePerson(newPerson models.Person) (int, error) {
	person, err := s.repo.GetPerson(newPerson)
	if err != nil {
		return -1, err
	}

	if person.Id != 0 {
		return -1, fmt.Errorf("such a person has already been created")
	}

	newPerson, err = encodingData(newPerson)
	if err != nil {
		return -1, err // 500
	}

	return s.repo.CreatePerson(newPerson)
}

func encodingData(person models.Person) (models.Person, error) {
	age, err := pkg.EnrichmentOfDataOnAge(person.Name)
	if err != nil {
		return person, err
	}

	gender, err := pkg.EnrichingDataOnGender(person.Name)
	if err != nil {
		return person, err
	}
	fmt.Println(gender)
	nationalize, err := pkg.EnrichmentOfDataOnNationality(person.Name)
	if err != nil {
		return person, err
	}

	if gender != "" {
		person.Gender = gender
	}
	if age != 0 {
		person.Age = age
	}
	if nationalize != nil {
		person.Nationalize = nationalize
	}
	fmt.Println(person)
	return person, nil
}
