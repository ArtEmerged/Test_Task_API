package models

type Person struct {
	Id          uint8  `json:"-"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         uint8  `json:"age"`
	Gender      string `json:"gender"`
	Nationalize []string
}

