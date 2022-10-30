package services

import (
	"github.com/golobby/container/v3"
	"mime/multipart"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
)

type UserService struct {
	repository repositories.IUserRepository
}

func CreateUserService() *UserService { return &UserService{repository: getUserRepository()} }

func (u UserService) Create(dto *user.CreateUserDTO) (*user.ReadUserDTO, error) {
	return u.repository.Create(dto)
}

func (u UserService) GetAll() ([]*user.ReadUserDTO, error) {
	return u.repository.GetAll()
}

func (u UserService) GetById(userId uint) (*user.ReadUserDTO, error) {
	return u.repository.GetById(userId)
}

func (u UserService) Delete(userId uint) error {
	return u.repository.Delete(userId)
}

func (u UserService) Update(userId uint, dto *user.UpdateUserDTO) (*user.ReadUserDTO, error) {
	return u.repository.Update(userId, dto)
}

func getUserRepository() repositories.IUserRepository {
	var injector repositories.IUserRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}

func (u UserService) UploadProfilePicture(username string, profilePicture *multipart.FileHeader) (string, error) {
	return u.repository.UploadProfilePicture(username, profilePicture)
}
