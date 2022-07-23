package services

import (
	"showcaseme/domain/DTO/skill_category"
)

type ISkillCategoryService interface {
	Create(dto *skill_category.CreateSkillCategoryDTO) (*skill_category.ReadSkillCategoryDTO, error)
	GetAll() ([]*skill_category.ReadSkillCategoryDTO, error)
	GetById(id uint) (*skill_category.ReadSkillCategoryDTO, error)
	Delete(id uint) error
	Update(id uint, dto *skill_category.UpdateSkillCategoryDTO) (*skill_category.ReadSkillCategoryDTO, error)
}
