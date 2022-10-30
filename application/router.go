package application

import (
	"github.com/gofiber/fiber/v2"
	"showcaseme/application/routes"
)

func RegisterRoutes(router fiber.Router) {
	routes.RegisterAuthRoutes(router)
	routes.RegisterArticleRoutes(router)
	routes.RegisterCarouselItemCategoryRoutes(router)
	routes.RegisterProjectCategoryRoutes(router)
	routes.RegisterProjectRoutes(router)
	routes.RegisterResumeRoutes(router)
	routes.RegisterSkillCategoryCategoryRoutes(router)
	routes.RegisterSkillRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterUserWebsiteRoutes(router)
	routes.RegisterWorkExperienceRoutes(router)
}
