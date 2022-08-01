package services

import (
	"showcaseme/domain/DTO/project_category"
)

type IProjectCategoryService interface {
	Create(dto *project_category.CreateProjectCategoryDTO) (*project_category.ReadProjectCategoryDTO, error)
	GetAll() ([]*project_category.ReadProjectCategoryDTO, error)
    GetById(id uint) (*project_category.ReadProjectCategoryDTO, error)
    Delete(id uint) error
    Update(id uint, dto *project_category.UpdateProjectCategoryDTO) (*project_category.ReadProjectCategoryDTO, error)
}
