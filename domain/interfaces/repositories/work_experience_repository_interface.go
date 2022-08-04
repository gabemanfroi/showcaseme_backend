package repositories

import (
	"showcaseme/domain/DTO/work_experience"
)

type IWorkExperienceRepository interface {
	Create(dto *work_experience.CreateWorkExperienceDTO) (*work_experience.ReadWorkExperienceDTO, error)
	GetAll() ([]*work_experience.ReadWorkExperienceDTO, error)
    GetById(id uint) (*work_experience.ReadWorkExperienceDTO, error)
    Delete(id uint) error
    Update(id uint, dto *work_experience.UpdateWorkExperienceDTO) (*work_experience.ReadWorkExperienceDTO, error)
}
