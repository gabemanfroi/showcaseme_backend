package repositories

import (
	"showcaseme/domain/DTO/auth"
	"showcaseme/domain/DTO/user"
)

type IAuthRepository interface {
	Login(dto *auth.LoginDTO) error
	Register(dto *auth.RegisterDTO) (*user.ReadUserDTO, error)
}
