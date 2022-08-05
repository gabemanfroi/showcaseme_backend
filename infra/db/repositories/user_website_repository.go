package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/user_website"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
)

type UserWebsiteRepository struct {
	sqlClient *gorm.DB
}

func CreateUserWebsiteRepository() *UserWebsiteRepository {
	return &UserWebsiteRepository{sqlClient: db.GetSqlInstance()}
}

func (repository UserWebsiteRepository) Create(dto *user_website.CreateUserWebsiteDTO) (*user_website.ReadUserWebsiteDTO, error) {
	u := models.UserWebsite{
		Url:    dto.Url,
		UserId: dto.UserId,
		Type:   dto.Type,
	}
	repository.sqlClient.Create(&u)

	if u.ID == 0 {
		return nil, errors.New("an error has occured when creating your user_website, verify")
	}

	createduserWebsite, _ := repository.GetById(u.ID)

	return createduserWebsite, nil
}

func (repository UserWebsiteRepository) GetAll() ([]*user_website.ReadUserWebsiteDTO, error) {
	var userWebsites []*models.UserWebsite
	var userWebsiteDTOs []*user_website.ReadUserWebsiteDTO

	repository.sqlClient.Find(&userWebsites)

	for _, u := range userWebsites {
		userWebsiteDTOs = append(userWebsiteDTOs, &user_website.ReadUserWebsiteDTO{
			ID:   u.ID,
			Url:  u.Url,
			Type: u.Type,
		})
	}

	return userWebsiteDTOs, nil
}

func (repository UserWebsiteRepository) GetById(id uint) (*user_website.ReadUserWebsiteDTO, error) {
	var u *models.UserWebsite

	repository.sqlClient.Find(&u, id)

	if u.ID == 0 {
		return nil, errors.New("user_website not found")
	}

	return &user_website.ReadUserWebsiteDTO{
		ID:   u.ID,
		Url:  u.Url,
		Type: u.Type,
	}, nil
}

func (repository UserWebsiteRepository) Delete(id uint) error {
	var u models.UserWebsite
	repository.sqlClient.Find(&u, id)
	if u.ID == 0 {
		return errors.New("user_website not found")
	}
	repository.sqlClient.Delete(&u)
	return nil
}

func (repository UserWebsiteRepository) Update(id uint, dto *user_website.UpdateUserWebsiteDTO) (*user_website.ReadUserWebsiteDTO, error) {
	var u models.UserWebsite

	repository.sqlClient.Find(&u, id)

	if u.ID == 0 {
		return nil, errors.New("user_website not found")
	}

	utils.UpdateModelValuesFromDTO(&u, dto)
	repository.sqlClient.Save(&u)

	updatedUserWebsite, _ := repository.GetById(u.ID)

	return updatedUserWebsite, nil
}
