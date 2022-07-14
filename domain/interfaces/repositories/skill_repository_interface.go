package repositories

import (
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/models"
)

type ISkillRepository interface {
	Create(dto *skill.CreateSkillDTO) *models.Skill
	GetAll() ([]*models.Skill, error)
	GetById(id uint) (*models.Skill, error)
	Delete(id uint) error
	Update(id uint, dto *skill.UpdateSkillDTO) (*models.Skill, error)
}
