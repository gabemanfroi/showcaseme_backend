# Services

---

<p>Expected to contain the implementations for the Services Interfaces located in <b>domain/interfaces/services</b></p>
<p>e.g</p>

```go

type UserService struct {
	repository repositories.IUserRepository
}

func CreateUserService() *UserService { return &UserService{repository: getUserRepository()} }

func (u UserService) Create(dto *user.CreateUserDTO) models.User {
	return u.repository.Create(dto)
}

func (u UserService) GetAll() ([]models.User, error) {
	return u.repository.GetAll()
}

func (u UserService) GetById(userid uint) (models.User, error) {
	return u.repository.GetById(userId)
}

func (u UserService) Delete(userid uint) error {
	return u.repository.Delete(userId)
}

func (u UserService) Update(userid uint, dto *user.UpdateUserDTO) (models.User, error) {
	return u.repository.Update(userId, dto)
}

func getUserRepository() repositories.IUserRepository {
	var injector repositories.IUserRepository
	utils.Check(container.Resolve(&injector), "Error while retrieving UserRepository instance")
	return injector
}
```