package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Proficiency uint8  `gorm:"not null"`
	UserId      uint8  `gorm:"not null"`
	User        User
	Active      bool
}
