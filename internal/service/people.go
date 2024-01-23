package service

import (
	"strings"
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

func (s *PeopleService) GetPersonById(id int) (models.Person, error) {
	return s.repo.GetPersonById(id)
}

func (s *PeopleService) GetPeople(filters models.Filters) ([]models.Person, error) {
	peopleId, err := s.repo.GetPeopleId(filters)
	if err != nil {
		return nil, err
	}

	people := make([]models.Person, len(peopleId))
	var person models.Person
	for ind, id := range peopleId {
		person, err = s.repo.GetPersonById(id)
		if err != nil {
			return nil, err
		}
		people[ind] = person
	}

	return people, nil
}

func (s *PeopleService) UpdatePerson(id int, person models.Person) error {
	oldPerson, err := s.repo.GetPersonById(id)
	if err != nil {
		return err
	}

	transferToUppercase(&person.Name, &person.Surname, &person.Patronymic)

	if person.Name != "" && person.Name != oldPerson.Name {
		person, err = encodingData(person)
		if err != nil {
			return err // 500
		}
	}

	return s.repo.UpdatePerson(id, person)
}

func (s *PeopleService) DeletePerson(id int) error {
	return s.repo.DeletePerson(id)
}

func (s *PeopleService) CreatePerson(newPerson models.Person) (int, error) {
	person, err := s.repo.GetPerson(newPerson)
	if err != nil {
		return -1, err
	}

	if person.Id != 0 {
		return -1, models.ErrAlreadyCreated
	}

	transferToUppercase(&newPerson.Name, &newPerson.Surname, &newPerson.Patronymic)

	newPerson, err = encodingData(newPerson)
	if err != nil {
		return -1, err // 500
	}

	return s.repo.CreatePerson(newPerson)
}

func transferToUppercase(letters ...*string) {
	for i := 0; i < len(letters); i++ {
		if len(*letters[i]) != 0 {
			newWord := []rune(strings.ToLower(*letters[i]))
			newWord[0] = newWord[0] - 32
			*letters[i] = string(newWord)
		}
	}
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

	return person, nil
}
