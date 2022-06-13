package repositories

import (
	"showcaseme/domain/models"
)

type SkillRepositoryInterface interface {
	Create(user *models.Skill) models.Skill
	GetAll() ([]models.Skill, error)
	GetById(id string) (models.User, error)
	Delete(id string) error
}
