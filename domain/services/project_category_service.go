package services

import (
	"showcaseme/domain/DTO/project_category"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
	"github.com/golobby/container/v3"
)

type ProjectCategoryService struct {
	repository repositories.IProjectCategoryRepository
}

func CreateProjectCategoryService() *ProjectCategoryService { return &ProjectCategoryService{repository: getProjectCategoryService()} }

func (service ProjectCategoryService) Create(dto *project_category.CreateProjectCategoryDTO) (*project_category.ReadProjectCategoryDTO, error) {
	return service.repository.Create(dto)
}

func (service ProjectCategoryService) GetAll() ([]*project_category.ReadProjectCategoryDTO, error) {
	return service.repository.GetAll()
}

func (service ProjectCategoryService) GetById(id uint) (*project_category.ReadProjectCategoryDTO, error) {
	return service.repository.GetById(id)
}

func (service ProjectCategoryService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service ProjectCategoryService) Update(id uint, dto *project_category.UpdateProjectCategoryDTO) (*project_category.ReadProjectCategoryDTO, error) {
	return service.repository.Update(id, dto)
}

func getProjectCategoryService() repositories.IProjectCategoryRepository {
	var injector repositories.IProjectCategoryRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving ProjectCategoryRepository instance")
	return injector
}
