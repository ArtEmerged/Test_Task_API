package handler

import (
	"fmt"
	"strings"
	"test_task/internal/models"
)

func validateData(person models.Person) error {
	data := fmt.Sprintf("%s%s%s", person.Name, person.Surname, person.Patronymic)

	if len(data) == 0 {
		return models.ErrInvalidData
	}

	data = strings.ToLower(data)
	for _, v := range data {
		if v < 'a' || v > 'z' {
			return models.ErrInvalidData
		}
	}
	return nil
}
