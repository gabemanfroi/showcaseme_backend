package services

import "showcaseme/domain/models"

type UserService interface {
	SaveUser(user *models.User)
	GetUsers() ([]models.User, error)
	GetUser(uint64) (*models.User, error)
}
