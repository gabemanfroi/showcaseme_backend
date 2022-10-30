package repositories

import (
	"errors"
	"gorm.io/gorm"
	"showcaseme/domain/DTO/auth"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/models"
	"showcaseme/infra/db"
)

type AuthRepository struct {
	sqlClient *gorm.DB
}

func (repository AuthRepository) Login(dto *auth.LoginDTO) error {
	var user models.User
	repository.sqlClient.Where(&models.User{Email: dto.Email, Password: dto.Password}).First(&user)

	if user.ID == 0 {
		return errors.New("wrong email or password")
	}

	return nil
}

func (repository AuthRepository) Register(dto *auth.RegisterDTO) (*user.ReadUserDTO, error) {
	var u models.User
	repository.sqlClient.Where(&models.User{Email: dto.Email}).First(&u)
	if u.ID != 0 {
		return nil, errors.New("email is already taken")
	}
	repository.sqlClient.Where(&models.User{Username: dto.Username}).First(&u)
	if u.ID != 0 {
		return nil, errors.New("username is already taken")
	}

	u = models.User{
		Email:    dto.Email,
		Password: dto.Password,
		Username: dto.Username,
	}
	repository.sqlClient.Create(&u)

	return &user.ReadUserDTO{ID: u.ID, Username: u.Username, Email: u.Email}, nil
}

func CreateAuthRepository() *AuthRepository {
	return &AuthRepository{sqlClient: db.GetSqlInstance()}
}
