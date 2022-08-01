package services

import (
	"showcaseme/domain/DTO/user_website"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
	"github.com/golobby/container/v3"
)

type UserWebsiteService struct {
	repository repositories.IUserWebsiteRepository
}

func CreateUserWebsiteService() *UserWebsiteService { return &UserWebsiteService{repository: getUserWebsiteService()} }

func (service UserWebsiteService) Create(dto *user_website.CreateUserWebsiteDTO) (*user_website.ReadUserWebsiteDTO, error) {
	return service.repository.Create(dto)
}

func (service UserWebsiteService) GetAll() ([]*user_website.ReadUserWebsiteDTO, error) {
	return service.repository.GetAll()
}

func (service UserWebsiteService) GetById(id uint) (*user_website.ReadUserWebsiteDTO, error) {
	return service.repository.GetById(id)
}

func (service UserWebsiteService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service UserWebsiteService) Update(id uint, dto *user_website.UpdateUserWebsiteDTO) (*user_website.ReadUserWebsiteDTO, error) {
	return service.repository.Update(id, dto)
}

func getUserWebsiteService() repositories.IUserWebsiteRepository {
	var injector repositories.IUserWebsiteRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserWebsiteRepository instance")
	return injector
}
