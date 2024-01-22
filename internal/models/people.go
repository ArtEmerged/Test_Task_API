package models

type Nationalize struct {
}

type Person struct {
	Id          uint8  `json:"-"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Patronymic  string `json:"patronymic"`
	Age         uint8  `json:"age"`
	Gender      string `json:"gender"`
	Nationalize []string
}
