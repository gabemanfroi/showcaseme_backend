package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	User    User
	UserId  uint   `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
