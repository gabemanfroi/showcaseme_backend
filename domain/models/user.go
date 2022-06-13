package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Age       int
	City      string
	Country   string
	Email     string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Pronouns  string
	Active    bool `gorm:"default:true"`
}
