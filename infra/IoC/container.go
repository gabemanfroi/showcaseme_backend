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
	utils.Check(container.Transient(func() repositories.IUserRepository { return repositoriesImpl.CreateUserRepository() }),
		"error while creating container bindings [Repositories - User]")
	utils.Check(container.Transient(func() repositories.ISkillRepository { return repositoriesImpl.CreateSkillRepository() }),
		"error while creating container bindings [Repositories - Skill]")
}

func bindServices() {
	utils.Check(container.Transient(func() services.IUserService { return servicesImpl.CreateUserService() }),
		"error while creating container bindings [Services - User]")
	utils.Check(container.Transient(func() services.ISkillService { return servicesImpl.CreateSkillService() }),
		"error while creating container bindings [Services - Skill]")
}

func bindControllers() {
	utils.Check(container.Transient(func() controllers.IUserController { return controllersImpl.CreateUserController() }),
		"error while creating container bindings [Controllers - User]")
	utils.Check(container.Transient(func() controllers.ISkillController { return controllersImpl.CreateSkillController() }),
		"error while creating container bindings [Controllers - User]")
}
