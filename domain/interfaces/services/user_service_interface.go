package services

import (
	"mime/multipart"
	"showcaseme/domain/DTO/user"
)

type IUserService interface {
	Create(dto *user.CreateUserDTO) (*user.ReadUserDTO, error)
	GetAll() ([]*user.ReadUserDTO, error)
	GetById(id uint) (*user.ReadUserDTO, error)
	Delete(id uint) error
	Update(id uint, dto *user.UpdateUserDTO) (*user.ReadUserDTO, error)
	UploadProfilePicture(username string, profilePicture *multipart.FileHeader) (string, error)
}
