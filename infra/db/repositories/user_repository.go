package repositories

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gorm.io/gorm"
	"mime/multipart"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/models"
	"showcaseme/infra"
	"showcaseme/infra/core"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

type UserRepository struct {
	sqlClient  *gorm.DB
	awsSession *session.Session
}

func CreateUserRepository() *UserRepository {
	return &UserRepository{
		sqlClient:  db.GetSqlInstance(),
		awsSession: infra.CreateAwsSession(),
	}
}

func (repository UserRepository) Create(dto *user.CreateUserDTO) (*user.ReadUserDTO, error) {
	var u models.User
	repository.sqlClient.Where(&models.User{Email: dto.Email}).First(&u)
	if u.ID != 0 {
		return nil, errors.New("email is already taken")
	}
	repository.sqlClient.Where(&models.User{Username: dto.Username}).First(&u)
	if u.ID != 0 {
		return nil, errors.New("username is already taken")
	}

	u = models.User{
		Age:       dto.Age,
		City:      dto.City,
		Country:   dto.Country,
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Pronouns:  dto.Pronouns,
		Password:  dto.Password,
		Username:  dto.Username,
		Role:      dto.Role,
	}
	repository.sqlClient.Create(&u)

	return &user.ReadUserDTO{ID: u.ID, Username: u.Username, Email: u.Email, FirstName: u.FirstName, LastName: u.LastName}, nil
}

func (repository UserRepository) GetAll() ([]*user.ReadUserDTO, error) {
	var users []*models.User
	var userDTOs []*user.ReadUserDTO

	repository.sqlClient.Find(&users)

	for _, u := range users {
		userDTOs = append(userDTOs, &user.ReadUserDTO{
			ID:                u.ID,
			Username:          u.Username,
			FirstName:         u.FirstName,
			LastName:          u.LastName,
			Email:             u.Email,
			Role:              u.Role,
			ProfilePictureUrl: u.ProfilePictureURL,
		})
	}

	return userDTOs, nil
}

func (repository UserRepository) GetById(id uint) (*user.ReadUserDTO, error) {
	var u models.User

	repository.sqlClient.Find(&u, id)

	if u.ID == 0 {
		return nil, errors.New("user not found")
	}
	return &user.ReadUserDTO{
		ID:                u.ID,
		Username:          u.Username,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		Email:             u.Email,
		Role:              u.Role,
		ProfilePictureUrl: u.ProfilePictureURL,
	}, nil
}

func (repository UserRepository) Delete(id uint) error {
	var u models.User

	repository.sqlClient.Find(&u, id)

	if u.ID == 0 {
		return errors.New("user not found")
	}

	repository.sqlClient.Where(&models.Skill{UserId: u.ID}).Delete(&models.Skill{})
	repository.sqlClient.Where(&models.SkillCategory{UserId: u.ID}).Delete(&models.SkillCategory{})
	repository.sqlClient.Where(&models.CarouselItem{UserId: u.ID}).Delete(&models.CarouselItem{})
	repository.sqlClient.Where(&models.UserWebsite{UserId: u.ID}).Delete(&models.UserWebsite{})
	repository.sqlClient.Where(&models.Article{UserId: u.ID}).Delete(&models.Article{})
	repository.sqlClient.Where(&models.ProjectCategory{UserId: u.ID}).Delete(&models.ProjectCategory{})
	repository.sqlClient.Where(&models.Project{UserId: u.ID}).Delete(&models.Project{})
	repository.sqlClient.Delete(&u)

	return nil
}

func (repository UserRepository) Update(id uint, dto *user.UpdateUserDTO) (*user.ReadUserDTO, error) {
	var u models.User

	repository.sqlClient.Find(&u, id)
	if u.ID == 0 {
		return nil, errors.New("user not found")
	}
	utils.UpdateModelValuesFromDTO(&u, dto)
	repository.sqlClient.Save(&u)

	return &user.ReadUserDTO{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		Username:  u.Username,
		LastName:  u.LastName,
	}, nil
}

func (repository UserRepository) UploadProfilePicture(username string, profilePicture *multipart.FileHeader) (string, error) {
	var u *models.User
	file, err := profilePicture.Open()

	fileType := profilePicture.Header["Content-Type"][0]

	uploader := s3manager.NewUploader(repository.awsSession)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(core.AppConfig.AwsBucketName),
		Key:         aws.String(profilePicture.Filename),
		Body:        file,
		ContentType: aws.String(fileType),
	})
	if err != nil {
		return "", err
	}

	filepath := "https://" + core.AppConfig.AwsBucketName + "." + "s3-" + core.AppConfig.AwsRegion + ".amazonaws.com/" + profilePicture.Filename
	repository.sqlClient.Where(&models.User{Username: username}).First(&u)
	u.ProfilePictureURL = filepath
	repository.sqlClient.Save(&u)

	return filepath, nil
}
