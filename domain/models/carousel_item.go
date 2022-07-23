package models

import "gorm.io/gorm"

type CarouselItem struct {
	gorm.Model
	Position uint   `gorm:"not null"`
	Content  string `gorm:"not null"`
	UserId   uint   `gorm:"not null"`
	User     User
}
