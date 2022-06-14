package services

import (
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/models"
)

type ISkillService interface {
	Create(dto *skill.CreateSkillDTO) models.Skill
	GetAll() ([]models.Skill, error)
	GetById(id string) (models.Skill, error)
	Delete(id string) error
	Update(id string, dto *skill.UpdateSkillDTO) (models.Skill, error)
}
