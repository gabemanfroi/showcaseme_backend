package IoC

import (
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
	controllersImpl "showcaseme/application/controllers"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/domain/interfaces/repositories"
	"showcaseme/domain/interfaces/services"
	servicesImpl "showcaseme/domain/services"
	"showcaseme/infra/db"
	repositoriesImpl "showcaseme/infra/db/repositories"
	"showcaseme/internal/utils"
)

type Constructors struct {
	controllersImpl.UserController
}

func InitContainer() {
	bindCore()
	bindRepositories()
	bindServices()
	bindControllers()
}

func bindCore() {
	err := container.Singleton(func() *gorm.DB { return db.CreateSqlInstance() })
	utils.Check(err, "error while creating container bindings [database]")
}

func bindRepositories() {
	err := container.Transient(func() repositories.UserRepositoryInterface { return repositoriesImpl.CreateUserRepository() })
	utils.Check(err, "error while creating container bindings [repositories]")
}

func bindServices() {
	err := container.Transient(func() services.UserServiceInterface { return servicesImpl.CreateUserService() })
	utils.Check(err, "error while creating container bindings [services]")
}

func bindControllers() {
	err := container.Transient(func() controllers.UserControllerInterface { return controllersImpl.CreateUserController() })
	utils.Check(err, "error while creating container bindings [controllers]")
}
