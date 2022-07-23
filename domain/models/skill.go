package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name            string `gorm:"not null"`
	Proficiency     uint8  `gorm:"not null"`
	UserId          uint   `gorm:"not null"`
	User            User
	Show            bool
	SkillCategoryId uint `gorm:"not null"`
	SkillCategory   SkillCategory
}
