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

func (repository UserRepository) Create(dto *user.CreateUserDTO) *models.User {
	u := models.User{
		Age:       dto.Age,
		City:      dto.City,
		Country:   dto.Country,
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Pronouns:  dto.Pronouns,
		Active:    true,
	}
	repository.sqlClient.Create(&u)
	return &u
}

func (repository UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	repository.sqlClient.Find(&users)
	return users, nil
}

func (repository UserRepository) GetById(id uint) (*models.User, error) {
	var u models.User
	repository.sqlClient.Find(&u, id)
	if u.ID == 0 {
		return &u, errors.New("user not found")
	}
	return &u, nil
}

func (repository UserRepository) Delete(id uint) error {
	var u models.User
	repository.sqlClient.Find(&u, id)
	if u.ID == 0 {
		return errors.New("user not found")
	}
	u.Active = false
	repository.sqlClient.Save(&u)
	return nil
}

func (repository UserRepository) Update(id uint, dto *user.UpdateUserDTO) (*models.User, error) {
	var u models.User
	repository.sqlClient.Find(&u, id)
	if u.ID == 0 {
		return &u, errors.New("user not found")
	}
	updateUserValuesFromDTO(&u, dto)
	repository.sqlClient.Save(&u)
	return &u, nil
}

func updateUserValuesFromDTO(model *models.User, dto *user.UpdateUserDTO) {
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
