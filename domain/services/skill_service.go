package services

import (
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/domain/models"
	"showcaseme/internal/utils"
)

type SkillService struct {
	repository repositories.ISkillRepository
}

func (service SkillService) Create(dto *skill.CreateSkillDTO) models.Skill {
	return service.repository.Create(dto)
}

func (service SkillService) GetAll() ([]models.Skill, error) {
	return service.repository.GetAll()
}

func (service SkillService) GetById(id string) (models.Skill, error) {
	return service.repository.GetById(id)
}

func (service SkillService) Delete(id string) error {
	return service.repository.Delete(id)
}

func (service SkillService) Update(id string, dto *skill.UpdateSkillDTO) (models.Skill, error) {
	return service.repository.Update(id, dto)
}

func CreateSkillService() *SkillService { return &SkillService{repository: getSkillRepository()} }

func getSkillRepository() repositories.ISkillRepository {
	var injector repositories.ISkillRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}
