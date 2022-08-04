package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"showcaseme/infra/db/repositories"
	"showcaseme/infra/tests/mocks"
)

var repository *repositories.UserRepository

func (s *RepositoriesTestSuite) TestGetAllActiveUsers() {
	users, err := repositories.CreateUserRepository().GetAll()
	require.NoError(s.T(), err)
	assert.Greater(s.T(), len(users), 0)
}

func (s *RepositoriesTestSuite) TestGetNonActiveUsers() {
	userToBeInactivated := mocks.CreateInactiveUser()
	createdUser, err := repositories.CreateUserRepository().Create(userToBeInactivated)
	require.NoError(s.T(), err)
	initialUsers, err := repositories.CreateUserRepository().GetAll()
	require.NoError(s.T(), err)

	repositories.CreateUserRepository().Delete(createdUser.ID)

	finalUsers, err := repositories.CreateUserRepository().GetAll()
	require.NoError(s.T(), err)

	assert.Greater(s.T(), len(initialUsers), len(finalUsers))
}
