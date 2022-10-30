package services

import (
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/auth"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/internal/utils"
)

type AuthService struct {
	repository repositories.IAuthRepository
}

func (service AuthService) Login(dto *auth.LoginDTO) error {
	return service.repository.Login(dto)
}

func (service AuthService) Register(dto *auth.RegisterDTO) (*user.ReadUserDTO, error) {
	return service.repository.Register(dto)
}

func CreateAuthService() *AuthService {
	return &AuthService{repository: getAuthRepository()}
}

func getAuthRepository() repositories.IAuthRepository {
	var injector repositories.IAuthRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving ArticleRepository instance")
	return injector
}
