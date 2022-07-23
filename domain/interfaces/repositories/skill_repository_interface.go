package repositories

import (
	"showcaseme/domain/DTO/skill"
)

type ISkillRepository interface {
	Create(dto *skill.CreateSkillDTO) (*skill.ReadSkillDTO, error)
	GetAll() ([]*skill.ReadSkillDTO, error)
	GetById(id uint) (*skill.ReadSkillDTO, error)
	Delete(id uint) error
	Update(id uint, dto *skill.UpdateSkillDTO) (*skill.ReadSkillDTO, error)
}
