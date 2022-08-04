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
	utils.Check(container.Transient(func() repositories.IUserWebsiteRepository { return repositoriesImpl.CreateUserWebsiteRepository() }),
		"error while creating container bindings [Repositories - UserWebsite]")
	utils.Check(container.Transient(func() repositories.IArticleRepository { return repositoriesImpl.CreateArticleRepository() }),
		"error while creating container bindings [Repositories - Article]")
	utils.Check(container.Transient(func() repositories.IProjectCategoryRepository {
		return repositoriesImpl.CreateProjectCategoryRepository()
	}), "error while creating container bindings [Repositories - ProjectCategory]")
	utils.Check(container.Transient(func() repositories.IProjectRepository {
		return repositoriesImpl.CreateProjectRepository()
	}), "error while creating container bindings [Repositories - Project]")
	utils.Check(container.Transient(func() repositories.IWorkExperienceRepository {
		return repositoriesImpl.CreateWorkExperienceRepository()
	}), "error while creating container bindings [Repositories - Project]")
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
	utils.Check(container.Transient(func() services.IUserWebsiteService { return servicesImpl.CreateUserWebsiteService() }),
		"error while creating container bindings [Services - UserWebsite]")
	utils.Check(container.Transient(func() services.IArticleService { return servicesImpl.CreateArticleService() }),
		"error while creating container bindings [Services - Article]")
	utils.Check(container.Transient(func() services.IProjectCategoryService { return servicesImpl.CreateProjectCategoryService() }),
		"error while creating container bindings [Services - ProjectCategory]")
	utils.Check(container.Transient(func() services.IProjectService { return servicesImpl.CreateProjectService() }),
		"error while creating container bindings [Services - Project]")
	utils.Check(container.Transient(func() services.IWorkExperienceService { return servicesImpl.CreateWorkExperienceService() }),
		"error while creating container bindings [Services - Project]")
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
	utils.Check(container.Transient(func() controllers.IUserWebsiteController { return controllersImpl.CreateUserWebsiteController() }),
		"error while creating container bindings [Controllers - UserWebsite]")
	utils.Check(container.Transient(func() controllers.IArticleController { return controllersImpl.CreateArticleController() }),
		"error while creating container bindings [Controllers - Article]")
	utils.Check(container.Transient(func() controllers.IProjectCategoryController {
		return controllersImpl.CreateProjectCategoryController()
	}),
		"error while creating container bindings [Controllers - ProjectCategory]")
	utils.Check(container.Transient(func() controllers.IProjectController {
		return controllersImpl.CreateProjectController()
	}), "error while creating container bindings [Controllers - Project]")
	utils.Check(container.Transient(func() controllers.IWorkExperienceController {
		return controllersImpl.CreateWorkExperienceController()
	}), "error while creating container bindings [Controllers - Project]")
}
