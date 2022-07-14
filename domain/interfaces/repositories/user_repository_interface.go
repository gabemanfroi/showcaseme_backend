package repositories

import (
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/models"
)

type IUserRepository interface {
	Create(dto *user.CreateUserDTO) *models.User
	GetAll() ([]*models.User, error)
	GetById(id uint) (*models.User, error)
	Delete(id uint) error
	Update(id uint, dto *user.UpdateUserDTO) (*models.User, error)
}
