package models

import "gorm.io/gorm"

type ProjectCategory struct {
	gorm.Model
	User  User 
	UserId  uint `gorm:"not null"`
	Name  string `gorm:"not null"`
}
