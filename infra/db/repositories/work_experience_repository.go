package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/work_experience"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type WorkExperienceRepository struct {
	sqlClient *gorm.DB
}

func CreateWorkExperienceRepository() *WorkExperienceRepository {
	return &WorkExperienceRepository{sqlClient: db.GetSqlInstance()}
}

func (repository WorkExperienceRepository) Create(dto *work_experience.CreateWorkExperienceDTO) (*work_experience.ReadWorkExperienceDTO, error) {
	w := models.WorkExperience{
		UserId:      dto.UserId,
		Role:        dto.Role,
		CompanyName: dto.CompanyName,
		StartDate:   dto.StartDate,
	}
	repository.sqlClient.Create(&w)

	if w.ID == 0 {
		return nil, errors.New("an error has occured when creating your work_experience, verify")
	}

	createdworkExperience, _ := repository.GetById(w.ID)

	return createdworkExperience, nil
}

func (repository WorkExperienceRepository) GetAll() ([]*work_experience.ReadWorkExperienceDTO, error) {
	var workExperiences []*models.WorkExperience
	var workExperienceDTOs []*work_experience.ReadWorkExperienceDTO

	repository.sqlClient.Find(&workExperiences)

	for _, w := range workExperiences {
		workExperienceDTOs = append(workExperienceDTOs, &work_experience.ReadWorkExperienceDTO{
			ID:          w.ID,
			Role:        w.Role,
			CompanyName: w.CompanyName,
			StartDate:   w.StartDate,
			Description: w.Description,
			EndDate:     w.EndDate,
		})
	}

	return workExperienceDTOs, nil
}

func (repository WorkExperienceRepository) GetById(id uint) (*work_experience.ReadWorkExperienceDTO, error) {
	var w *models.WorkExperience

	repository.sqlClient.Find(&w, id)

	if w.ID == 0 {
		return nil, errors.New("work_experience not found")
	}

	return &work_experience.ReadWorkExperienceDTO{
		ID:          w.ID,
		Role:        w.Role,
		CompanyName: w.CompanyName,
		StartDate:   w.StartDate,
		Description: w.Description,
		EndDate:     w.EndDate,
	}, nil
}

func (repository WorkExperienceRepository) Delete(id uint) error {
	var w models.WorkExperience
	repository.sqlClient.Find(&w, id)
	if w.ID == 0 {
		return errors.New("work_experience not found")
	}
	repository.sqlClient.Delete(&w)
	return nil
}

func (repository WorkExperienceRepository) Update(id uint, dto *work_experience.UpdateWorkExperienceDTO) (*work_experience.ReadWorkExperienceDTO, error) {
	var w models.WorkExperience

	repository.sqlClient.Find(&w, id)

	if w.ID == 0 {
		return nil, errors.New("work_experience not found")
	}

	updateWorkExperienceValuesFromDTO(&w, dto)
	repository.sqlClient.Save(&w)

	updatedWorkExperience, _ := repository.GetById(w.ID)

	return updatedWorkExperience, nil
}

func updateWorkExperienceValuesFromDTO(model *models.WorkExperience, dto *work_experience.UpdateWorkExperienceDTO) {
	if dto.Role != nil {
		model.Role = *dto.Role
	}
	if dto.CompanyName != nil {
		model.CompanyName = *dto.CompanyName
	}
	if dto.EndDate != nil {
		model.EndDate = dto.EndDate
	}
	if dto.Description != nil {
		model.Description = *dto.Description
	}
	if dto.StartDate != nil {
		model.StartDate = *dto.StartDate
	}
}
