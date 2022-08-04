package models

import (
	"gorm.io/gorm"
	"time"
)

type WorkExperience struct {
	gorm.Model
	User        User
	UserId      uint      `gorm:"not null"`
	Role        string    `gorm:"not null"`
	CompanyName string    `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	Description string
	EndDate     *time.Time
}
