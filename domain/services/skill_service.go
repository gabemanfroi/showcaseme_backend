package services

import (
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
)

type SkillService struct {
	repository repositories.ISkillRepository
}

func CreateSkillService() *SkillService { return &SkillService{repository: getSkillRepository()} }

func (service SkillService) Create(dto *skill.CreateSkillDTO) (*skill.ReadSkillDTO, error) {
	return service.repository.Create(dto)
}

func (service SkillService) GetAll() ([]*skill.ReadSkillDTO, error) {
	return service.repository.GetAll()
}

func (service SkillService) GetById(id uint) (*skill.ReadSkillDTO, error) {
	return service.repository.GetById(id)
}

func (service SkillService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service SkillService) Update(id uint, dto *skill.UpdateSkillDTO) (*skill.ReadSkillDTO, error) {
	return service.repository.Update(id, dto)
}

func getSkillRepository() repositories.ISkillRepository {
	var injector repositories.ISkillRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}
