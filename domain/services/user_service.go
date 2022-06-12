package services

import "showcaseme/domain/models"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) SaveUser(user *models.User) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUsers() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUser(u2 uint64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
