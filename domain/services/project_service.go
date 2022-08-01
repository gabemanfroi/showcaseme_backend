package services

import (
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
	"github.com/golobby/container/v3"
)

type ProjectService struct {
	repository repositories.IProjectRepository
}

func CreateProjectService() *ProjectService { return &ProjectService{repository: getProjectService()} }

func (service ProjectService) Create(dto *project.CreateProjectDTO) (*project.ReadProjectDTO, error) {
	return service.repository.Create(dto)
}

func (service ProjectService) GetAll() ([]*project.ReadProjectDTO, error) {
	return service.repository.GetAll()
}

func (service ProjectService) GetById(id uint) (*project.ReadProjectDTO, error) {
	return service.repository.GetById(id)
}

func (service ProjectService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service ProjectService) Update(id uint, dto *project.UpdateProjectDTO) (*project.ReadProjectDTO, error) {
	return service.repository.Update(id, dto)
}

func getProjectService() repositories.IProjectRepository {
	var injector repositories.IProjectRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving ProjectRepository instance")
	return injector
}
