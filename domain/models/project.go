package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	User              User            `gorm:"not null"`
	UserId            uint            `gorm:"not null"`
	ProjectCategory   ProjectCategory `gorm:"not null"`
	ProjectCategoryId uint            `gorm:"not null"`
	Title             string          `gorm:"not null"`
	Url               string          `gorm:"not null"`
	ImageUrl          string
}
