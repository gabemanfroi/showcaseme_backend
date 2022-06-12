package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string
	Proficiency int
	UserId      int
	User        User
}
