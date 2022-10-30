package repositories

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/DTO/project_category"
	"showcaseme/domain/models"
	"showcaseme/infra"
	"showcaseme/infra/core"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

type ProjectRepository struct {
	sqlClient  *gorm.DB
	awsSession *session.Session
}

func CreateProjectRepository() *ProjectRepository {
	return &ProjectRepository{sqlClient: db.GetSqlInstance(),
		awsSession: infra.CreateAwsSession()}
}

func (repository ProjectRepository) Create(dto *project.CreateProjectDTO) (*project.ReadProjectDTO, error) {

	file, err := dto.BackgroundImage.Open()

	fileType := dto.BackgroundImage.Header["Content-Type"][0]

	uploader := s3manager.NewUploader(repository.awsSession)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(core.AppConfig.AwsBucketName),
		Key:         aws.String(dto.BackgroundImage.Filename),
		Body:        file,
		ContentType: aws.String(fileType),
	})

	if err != nil {
		return nil, err
	}

	filepath := "https://" + core.AppConfig.AwsBucketName + "." + "s3-" + core.AppConfig.AwsRegion + ".amazonaws.com/" + dto.BackgroundImage.Filename

	p := models.Project{
		UserId:            dto.UserId,
		ProjectCategoryId: dto.ProjectCategoryId,
		Title:             dto.Title,
		Url:               dto.Url,
		ImageUrl:          filepath,
	}
	repository.sqlClient.Create(&p)

	if p.ID == 0 {
		return nil, errors.New("an error has occured when creating your project, verify")
	}

	createdProject, _ := repository.GetById(p.ID)

	return createdProject, nil
}

func (repository ProjectRepository) GetAll() ([]*project.ReadProjectDTO, error) {
	var projects []*models.Project
	var projectDTOs []*project.ReadProjectDTO

	repository.sqlClient.Joins("ProjectCategory").Find(&projects)

	for _, p := range projects {
		projectDTOs = append(projectDTOs, &project.ReadProjectDTO{
			ID:              p.ID,
			ProjectCategory: &project_category.ReadProjectCategoryDTO{ID: p.ProjectCategoryId, Name: p.ProjectCategory.Name},
			ImageUrl:        p.ImageUrl,
			Url:             p.Url,
			Title:           p.Title,
		})
	}

	return projectDTOs, nil
}

func (repository ProjectRepository) GetById(id uint) (*project.ReadProjectDTO, error) {
	var p *models.Project

	repository.sqlClient.Joins("ProjectCategory").Find(&p, id)

	if p.ID == 0 {
		return nil, errors.New("project not found")
	}

	return &project.ReadProjectDTO{
		ID:              p.ID,
		ProjectCategory: &project_category.ReadProjectCategoryDTO{ID: p.ProjectCategoryId, Name: p.ProjectCategory.Name},
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

	utils.UpdateModelValuesFromDTO(&p, dto)
	repository.sqlClient.Save(&p)

	updatedProject, _ := repository.GetById(p.ID)

	return updatedProject, nil
}
