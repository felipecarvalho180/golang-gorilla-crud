package models

type Person struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Persons []Person

func (Person) TableName() string {
	return "persons"
}
