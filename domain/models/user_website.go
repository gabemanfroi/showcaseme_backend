package models

import "gorm.io/gorm"

type UserWebsite struct {
	gorm.Model
	Url    string `gorm:"not null"`
	Type   string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
	User   User
}
