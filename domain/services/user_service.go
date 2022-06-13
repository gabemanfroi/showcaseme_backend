package services

import (
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/domain/models"
	"showcaseme/internal/utils"
)

type UserService struct {
	repository repositories.UserRepositoryInterface
}

func CreateUserService() *UserService { return &UserService{repository: getRepository()} }

func (u UserService) Create(dto *user.CreateUserDTO) models.User {
	return u.repository.Create(dto)
}

func (u UserService) GetAll() ([]models.User, error) {
	return u.repository.GetAll()
}

func (u UserService) GetById(userId string) (models.User, error) {
	return u.repository.GetById(userId)
}

func (u UserService) Delete(userId string) error {
	return u.repository.Delete(userId)
}

func (u UserService) Update(userId string, dto *user.UpdateUserDTO) (models.User, error) {
	return u.repository.Update(userId, dto)
}

func getRepository() repositories.UserRepositoryInterface {
	var injector repositories.UserRepositoryInterface
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}
