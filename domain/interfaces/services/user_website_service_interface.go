package services

import (
	"showcaseme/domain/DTO/user_website"
)

type IUserWebsiteService interface {
	Create(dto *user_website.CreateUserWebsiteDTO) (*user_website.ReadUserWebsiteDTO, error)
	GetAll() ([]*user_website.ReadUserWebsiteDTO, error)
    GetById(id uint) (*user_website.ReadUserWebsiteDTO, error)
    Delete(id uint) error
    Update(id uint, dto *user_website.UpdateUserWebsiteDTO) (*user_website.ReadUserWebsiteDTO, error)
}
