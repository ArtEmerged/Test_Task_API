package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"test_task/internal/models"
)

type PeopleRepo struct {
	db *sql.DB
}

func NewPeoplePostgres(db *sql.DB) *PeopleRepo {
	return &PeopleRepo{db: db}
}

func (r *PeopleRepo) GetPersonById(id int) (models.Person, error) {
	var person models.Person
	query := fmt.Sprintf("SELECT * FROM  %s WHERE id=$1", peopleTable)
	err := r.db.QueryRow(query, id).Scan(
		&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			err = models.ErrNoSuchPerson
		}
		return person, err
	}
	query = fmt.Sprintf("SELECT country_id FROM %s WHERE person_id=$1", nationalizeTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return person, err
	}
	defer rows.Close()
	for rows.Next() {
		var country_id string
		err := rows.Scan(&country_id)
		if err != nil {
			return person, err
		}
		person.Nationalize = append(person.Nationalize, country_id)
	}
	err = rows.Err()
	if err != nil {
		return person, err
	}

	return person, nil
}

func (r *PeopleRepo) GetPeopleId(filters models.Filters) ([]int, error) {
	var peopleId []int

	query := fmt.Sprintf("SELECT p.id FROM %s p", peopleTable)
	var args []interface{}
	where := " WHERE 1=1"

	if nation := filters.Filters["nationalize"]; nation != "" {
		fmt.Println("nationalize == ok")
		args = append(args, nation)
		query += fmt.Sprintf(" JOIN %s n ON p.id = n.person_id%s", nationalizeTable, where)
		query += fmt.Sprintf(" AND n.country_id = $%d", len(args))
	} else {
		query += where
	}

	for key, value := range filters.Filters {
		if key != "nationalize" && value != "" {
			fmt.Println(key)
			args = append(args, value)
			query += fmt.Sprintf(" AND p.%s = $%d", key, len(args))
		}
	}

	query += fmt.Sprintf(" OFFSET %d LIMIT %d", filters.Offset, filters.Limit)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		peopleId = append(peopleId, id)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return peopleId, nil
}

func (r *PeopleRepo) UpdatePerson(id int, person models.Person) error {
	set := make([]string, 0)
	args := make([]interface{}, 0)

	if person.Name != "" {
		args = append(args, person.Name)
		set = append(set, fmt.Sprintf("name = $%d", len(args)))
	}
	if person.Surname != "" {
		args = append(args, person.Surname)
		set = append(set, fmt.Sprintf("surname = $%d", len(args)))
	}
	if person.Patronymic != "" {
		args = append(args, person.Patronymic)
		set = append(set, fmt.Sprintf("patronymic = $%d", len(args)))
	}
	if person.Age != 0 {
		args = append(args, person.Age)
		set = append(set, fmt.Sprintf("age = $%d", len(args)))
	}
	if person.Gender != "" {
		args = append(args, person.Gender)
		set = append(set, fmt.Sprintf("gender = $%d", len(args)))
	}

	args = append(args, id)
	joinSet := strings.Join(set, ", ")
	query := fmt.Sprintf("UPDATE %s  SET %s WHERE id = $%d", peopleTable, joinSet, len(args))
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(person.Nationalize) != 0 {
		query = fmt.Sprintf("DELETE FROM %s WHERE person_id=$1", nationalizeTable)
		_, err = tx.Exec(query, id)
		if err != nil {
			tx.Rollback()
			return err
		}

		query = fmt.Sprintf("INSERT INTO %s (person_id, country_id) VALUES($1,$2)", nationalizeTable)
		for _, country_id := range person.Nationalize {
			_, err = tx.Exec(query, id, country_id)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit()
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
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	deleteId, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if deleteId == 0 {
		return models.ErrNoSuchPerson
	}

	return nil
}
