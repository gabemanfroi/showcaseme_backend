package models

import "gorm.io/gorm"

type SkillCategory struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserId uint
	User   User
}
