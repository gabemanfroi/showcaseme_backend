package repositories

import (
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/models"
)

type IUserRepository interface {
	Create(dto *user.CreateUserDTO) models.User
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
	Delete(id string) error
	Update(id string, dto *user.UpdateUserDTO) (models.User, error)
}
