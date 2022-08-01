package application

import (
	"github.com/gofiber/fiber/v2"
	"showcaseme/application/routes"
)

func RegisterRoutes(router fiber.Router) {
	routes.RegisterCarouselItemCategoryRoutes(router)
	routes.RegisterUserRoutes(router)
	routes.RegisterSkillCategoryCategoryRoutes(router)
	routes.RegisterSkillRoutes(router)
	routes.RegisterResumeRoutes(router)
	routes.RegisterUserWebsiteRoutes(router)
	routes.RegisterArticleRoutes(router)
	routes.RegisterProjectCategoryRoutes(router)
	routes.RegisterProjectRoutes(router)
}
