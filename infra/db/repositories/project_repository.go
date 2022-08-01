package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type ProjectRepository struct {
	sqlClient *gorm.DB
}

func CreateProjectRepository() *ProjectRepository {
	return &ProjectRepository{sqlClient: db.GetSqlInstance()}
}

func (repository ProjectRepository) Create(dto *project.CreateProjectDTO) (*project.ReadProjectDTO, error) {
	p := models.Project{
		UserId:            dto.UserId,
		ProjectCategoryId: dto.ProjectCategoryId,
		Title:             dto.Title,
		Url:               dto.Url,
	}
	repository.sqlClient.Create(&p)

	if p.ID == 0 {
		return nil, errors.New("an error has occured when creating your project, verify")
	}

	createdproject, _ := repository.GetById(p.ID)

	return createdproject, nil
}

func (repository ProjectRepository) GetAll() ([]*project.ReadProjectDTO, error) {
	var projects []*models.Project
	var projectDTOs []*project.ReadProjectDTO

	repository.sqlClient.Find(&projects)

	for _, p := range projects {
		projectDTOs = append(projectDTOs, &project.ReadProjectDTO{
			ProjectCategory: p.ProjectCategory,
			ImageUrl:        p.ImageUrl,
			Url:             p.Url,
			Title:           p.Title,
		})
	}

	return projectDTOs, nil
}

func (repository ProjectRepository) GetById(id uint) (*project.ReadProjectDTO, error) {
	var p *models.Project

	repository.sqlClient.Find(&p, id)

	if p.ID == 0 {
		return nil, errors.New("project not found")
	}

	return &project.ReadProjectDTO{
		ProjectCategory: p.ProjectCategory,
		ImageUrl:        p.ImageUrl,
		Url:             p.Url,
		Title:           p.Title,
	}, nil
}

func (repository ProjectRepository) Delete(id uint) error {
	var p models.Project
	repository.sqlClient.Find(&p, id)
	if p.ID == 0 {
		return errors.New("project not found")
	}
	repository.sqlClient.Delete(&p)
	return nil
}

func (repository ProjectRepository) Update(id uint, dto *project.UpdateProjectDTO) (*project.ReadProjectDTO, error) {
	var p models.Project

	repository.sqlClient.Find(&p, id)

	if p.ID == 0 {
		return nil, errors.New("project not found")
	}

	updateProjectValuesFromDTO(&p, dto)
	repository.sqlClient.Save(&p)

	updatedProject, _ := repository.GetById(p.ID)

	return updatedProject, nil
}

func updateProjectValuesFromDTO(model *models.Project, dto *project.UpdateProjectDTO) {
	if dto.ProjectCategoryId != nil {
		model.ProjectCategoryId = *dto.ProjectCategoryId
	}
	if dto.Title != nil {
		model.Title = *dto.Title
	}
	if dto.ImageUrl != nil {
		model.ImageUrl = *dto.ImageUrl
	}
	if dto.Url != nil {
		model.Url = *dto.Url
	}
}
