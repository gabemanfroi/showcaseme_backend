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
	utils.Check(container.Transient(func() repositories.ISkillCategoryRepository { return repositoriesImpl.CreateSkillCategoryRepository() }),
		"error while creating container bindings [Repositories - SkillCategory]")
	utils.Check(container.Transient(func() repositories.ISkillRepository { return repositoriesImpl.CreateSkillRepository() }),
		"error while creating container bindings [Repositories - Skill]")
	utils.Check(container.Transient(func() repositories.IResumeRepository { return repositoriesImpl.CreateResumeRepository() }),
		"error while creating container bindings [Repositories - Resume]")
	utils.Check(container.Transient(func() repositories.ICarouselItemRepository { return repositoriesImpl.CreateCarouselItemRepository() }),
		"error while creating container bindings [Repositories - CarouselItem]")
}

func bindServices() {
	utils.Check(container.Transient(func() services.IUserService { return servicesImpl.CreateUserService() }),
		"error while creating container bindings [Services - User]")
	utils.Check(container.Transient(func() services.ISkillCategoryService { return servicesImpl.CreateSkillCategoryService() }),
		"error while creating container bindings [Services - SkillCategory]")
	utils.Check(container.Transient(func() services.ISkillService { return servicesImpl.CreateSkillService() }),
		"error while creating container bindings [Services - Skill]")
	utils.Check(container.Transient(func() services.IResumeService { return servicesImpl.CreateResumeService() }),
		"error while creating container bindings [Services - Resume]")
	utils.Check(container.Transient(func() services.ICarouselItemService { return servicesImpl.CreateCarouselItemService() }),
		"error while creating container bindings [Services - CarouselItem]")
}

func bindControllers() {
	utils.Check(container.Transient(func() controllers.IUserController { return controllersImpl.CreateUserController() }),
		"error while creating container bindings [Controllers - User]")
	utils.Check(container.Transient(func() controllers.ISkillCategoryController { return controllersImpl.CreateSkillCategoryController() }),
		"error while creating container bindings [Controllers - SkillCategory]")
	utils.Check(container.Transient(func() controllers.ISkillController { return controllersImpl.CreateSkillController() }),
		"error while creating container bindings [Controllers - Skill]")
	utils.Check(container.Transient(func() controllers.IResumeController { return controllersImpl.CreateResumeController() }),
		"error while creating container bindings [Controllers - Resume]")
	utils.Check(container.Transient(func() controllers.ICarouselItemController { return controllersImpl.CreateCarouselItemController() }),
		"error while creating container bindings [Controllers - CarouselItem]")
}
