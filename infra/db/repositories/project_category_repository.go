package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/project_category"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

type ProjectCategoryRepository struct {
	sqlClient *gorm.DB
}

func CreateProjectCategoryRepository() *ProjectCategoryRepository {
	return &ProjectCategoryRepository{sqlClient: db.GetSqlInstance()}
}

func (repository ProjectCategoryRepository) Create(dto *project_category.CreateProjectCategoryDTO) (*project_category.ReadProjectCategoryDTO, error) {
	p := models.ProjectCategory{
		UserId: dto.UserId,
		Name:   dto.Name,
	}
	repository.sqlClient.Create(&p)

	if p.ID == 0 {
		return nil, errors.New("an error has occured when creating your project_category, verify")
	}

	createdprojectCategory, _ := repository.GetById(p.ID)

	return createdprojectCategory, nil
}

func (repository ProjectCategoryRepository) GetAll() ([]*project_category.ReadProjectCategoryDTO, error) {
	var projectCategorys []*models.ProjectCategory
	var projectCategoryDTOs []*project_category.ReadProjectCategoryDTO

	repository.sqlClient.Find(&projectCategorys)

	for _, p := range projectCategorys {
		projectCategoryDTOs = append(projectCategoryDTOs, &project_category.ReadProjectCategoryDTO{
			ID:   p.ID,
			Name: p.Name,
		})
	}

	return projectCategoryDTOs, nil
}

func (repository ProjectCategoryRepository) GetById(id uint) (*project_category.ReadProjectCategoryDTO, error) {
	var p *models.ProjectCategory

	repository.sqlClient.Find(&p, id)

	if p.ID == 0 {
		return nil, errors.New("project_category not found")
	}

	return &project_category.ReadProjectCategoryDTO{
		ID:   p.ID,
		Name: p.Name,
	}, nil
}

func (repository ProjectCategoryRepository) Delete(id uint) error {
	var p models.ProjectCategory
	repository.sqlClient.Find(&p, id)
	if p.ID == 0 {
		return errors.New("project_category not found")
	}

	repository.sqlClient.Where(&models.Project{ProjectCategoryId: p.ID}).Delete(&models.Project{})
	repository.sqlClient.Delete(&p)
	return nil
}

func (repository ProjectCategoryRepository) Update(id uint, dto *project_category.UpdateProjectCategoryDTO) (*project_category.ReadProjectCategoryDTO, error) {
	var p models.ProjectCategory

	repository.sqlClient.Find(&p, id)

	if p.ID == 0 {
		return nil, errors.New("project_category not found")
	}

	utils.UpdateModelValuesFromDTO(&p, dto)
	repository.sqlClient.Save(&p)

	updatedProjectCategory, _ := repository.GetById(p.ID)

	return updatedProjectCategory, nil
}
