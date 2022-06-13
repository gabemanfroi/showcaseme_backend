package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string `binding:required`
	Proficiency int    `binding:required`
	UserId      int    `binding:required`
	User        User
}
