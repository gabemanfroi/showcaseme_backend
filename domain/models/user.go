package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
	Country   string
	City      string
	Pronouns  string
	Age       int
}
