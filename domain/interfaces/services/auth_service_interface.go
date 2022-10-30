package services

import (
	"showcaseme/domain/DTO/auth"
	"showcaseme/domain/DTO/user"
)

type IAuthService interface {
	Login(dto *auth.LoginDTO) error
	Register(dto *auth.RegisterDTO) (*user.ReadUserDTO, error)
}
