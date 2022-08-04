package mocks

import (
	"github.com/brianvoe/gofakeit/v6"
	"showcaseme/domain/DTO/user"
)

func CreateInactiveUser() *user.CreateUserDTO {
	return &user.CreateUserDTO{
		Age:       gofakeit.Number(15, 60),
		City:      gofakeit.City(),
		Country:   gofakeit.Country(),
		Email:     gofakeit.Email(),
		FirstName: gofakeit.FirstName(),
		Username:  gofakeit.Username(),
		LastName:  gofakeit.LastName(),
		Pronouns:  gofakeit.Pronoun(),
		Password:  gofakeit.Password(true, true, false, true, false, 10),
		Role:      gofakeit.JobTitle(),
	}
}
