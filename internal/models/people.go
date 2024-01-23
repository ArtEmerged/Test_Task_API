package models

type Person struct {
	Id          uint8    `json:"id"`
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	Patronymic  string   `json:"patronymic"`
	Age         uint8    `json:"age"`
	Gender      string   `json:"gender"`
	Nationalize []string `json:"nationalize"`
}
type Filters struct {
	Offset  int
	Limit   int
	Filters map[string]interface{}
}
