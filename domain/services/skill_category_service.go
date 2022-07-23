package services

import (
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/skill_category"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
)

type SkillCategoryService struct {
	repository repositories.ISkillCategoryRepository
}

func CreateSkillCategoryService() *SkillCategoryService {
	return &SkillCategoryService{repository: getSkillCategoryRepository()}
}

func (service SkillCategoryService) Create(dto *skill_category.CreateSkillCategoryDTO) (*skill_category.ReadSkillCategoryDTO, error) {
	return service.repository.Create(dto)
}

func (service SkillCategoryService) GetAll() ([]*skill_category.ReadSkillCategoryDTO, error) {
	return service.repository.GetAll()
}

func (service SkillCategoryService) GetById(id uint) (*skill_category.ReadSkillCategoryDTO, error) {
	return service.repository.GetById(id)
}

func (service SkillCategoryService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service SkillCategoryService) Update(id uint, dto *skill_category.UpdateSkillCategoryDTO) (*skill_category.ReadSkillCategoryDTO, error) {
	return service.repository.Update(id, dto)
}

func getSkillCategoryRepository() repositories.ISkillCategoryRepository {
	var injector repositories.ISkillCategoryRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}
