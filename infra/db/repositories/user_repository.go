package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type UserRepository struct {
	sqlClient *gorm.DB
}

func CreateUserRepository() *UserRepository {
	return &UserRepository{sqlClient: db.GetSqlInstance()}
}

func (u UserRepository) Create(dto *user.CreateUserDTO) models.User {
	user := models.User{
		Age:       dto.Age,
		City:      dto.City,
		Country:   dto.Country,
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Pronouns:  dto.Pronouns,
		Active:    true,
	}

	u.sqlClient.Create(&user)

	return user
}

func (u UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	u.sqlClient.Find(&users)
	return users, nil
}

func (u UserRepository) GetById(id string) (models.User, error) {
	var user models.User
	u.sqlClient.Find(&user, id)
	return user, nil
}

func (u UserRepository) Delete(id string) error {
	var user models.User
	u.sqlClient.Find(&user, id)
	if user.ID == 0 {
		return errors.New("user not found")
	}
	user.Active = false
	u.sqlClient.Save(&user)
	return nil
}

func (u UserRepository) Update(id string, dto *user.UpdateUserDTO) (models.User, error) {
	var user models.User
	u.sqlClient.Find(&user, id)
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	updateValuesFromDto(&user, dto)
	u.sqlClient.Save(&user)
	return user, nil
}

func updateValuesFromDto(model *models.User, dto *user.UpdateUserDTO) {
	if dto.Age != nil {
		model.Age = *dto.Age
	}
	if dto.City != nil {
		model.City = *dto.City
	}
	if dto.Country != nil {
		model.Country = *dto.Country
	}
	if dto.Email != nil {
		model.Email = *dto.Email
	}
	if dto.FirstName != nil {
		model.FirstName = *dto.FirstName
	}
	if dto.LastName != nil {
		model.LastName = *dto.LastName
	}
	if dto.Pronouns != nil {
		model.Pronouns = *dto.Pronouns
	}
}
