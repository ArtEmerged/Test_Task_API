package repository

import (
	"database/sql"
	"fmt"
	"test_task/internal/models"
)

type PeopleRepo struct {
	db *sql.DB
}

func NewPeoplePostgres(db *sql.DB) *PeopleRepo {
	return &PeopleRepo{db: db}
}

func (r *PeopleRepo) GetPerson(newPerson models.Person) (models.Person, error) {
	var person models.Person
	query := fmt.Sprintf("SELECT * FROM  %s WHERE name=$1 AND surname=$2 AND patronymic=$3", peopleTable)
	err := r.db.QueryRow(query, newPerson.Name, newPerson.Surname, newPerson.Patronymic).Scan(
		&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender)
	if err != nil && err != sql.ErrNoRows {
		return person, err
	}

	return person, nil
}

func (r *PeopleRepo) CreatePerson(newPerson models.Person) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, age, gender) VALUES($1,$2,$3,$4,$5)RETURNING id", peopleTable)

	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	err = tx.QueryRow(query, newPerson.Name, newPerson.Surname, newPerson.Patronymic, newPerson.Age, newPerson.Gender).Scan(&id)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	if len(newPerson.Nationalize) == 0 {
		return id, tx.Commit()
	}

	query = fmt.Sprintf("INSERT INTO %s (person_id, country_id) VALUES($1,$2)", nationalizeTable)
	for _, country_id := range newPerson.Nationalize {
		_, err = tx.Exec(query, id, country_id)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	return id, tx.Commit()
}

func (r *PeopleRepo) DeletePerson(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", peopleTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNoSuchPerson
		}

		return err
	}

	return nil
}
