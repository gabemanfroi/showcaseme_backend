package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Age       int
	City      string
	Country   string
	Email     string `gorm:"not null,unique"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Username  string `gorm:"not null"`
	Pronouns  string
	Password  string `gorm:"not null"`
	Role      string
}
