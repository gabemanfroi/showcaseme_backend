package services

import (
	"showcaseme/domain/DTO/project"
)

type IProjectService interface {
	Create(dto *project.CreateProjectDTO) (*project.ReadProjectDTO, error)
	GetAll() ([]*project.ReadProjectDTO, error)
    GetById(id uint) (*project.ReadProjectDTO, error)
    Delete(id uint) error
    Update(id uint, dto *project.UpdateProjectDTO) (*project.ReadProjectDTO, error)
}
