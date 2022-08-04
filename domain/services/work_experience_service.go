package services

import (
	"showcaseme/domain/DTO/work_experience"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
	"github.com/golobby/container/v3"
)

type WorkExperienceService struct {
	repository repositories.IWorkExperienceRepository
}

func CreateWorkExperienceService() *WorkExperienceService { return &WorkExperienceService{repository: getWorkExperienceService()} }

func (service WorkExperienceService) Create(dto *work_experience.CreateWorkExperienceDTO) (*work_experience.ReadWorkExperienceDTO, error) {
	return service.repository.Create(dto)
}

func (service WorkExperienceService) GetAll() ([]*work_experience.ReadWorkExperienceDTO, error) {
	return service.repository.GetAll()
}

func (service WorkExperienceService) GetById(id uint) (*work_experience.ReadWorkExperienceDTO, error) {
	return service.repository.GetById(id)
}

func (service WorkExperienceService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service WorkExperienceService) Update(id uint, dto *work_experience.UpdateWorkExperienceDTO) (*work_experience.ReadWorkExperienceDTO, error) {
	return service.repository.Update(id, dto)
}

func getWorkExperienceService() repositories.IWorkExperienceRepository {
	var injector repositories.IWorkExperienceRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving WorkExperienceRepository instance")
	return injector
}
