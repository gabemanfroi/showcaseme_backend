package routes

import (
	schema_validation_middleware "github.com/gabemanfroi/schema-validation-middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/project_category"
	"showcaseme/internal/utils"
)

func RegisterProjectCategoryRoutes(router fiber.Router) {
	var controller controllers.IProjectCategoryController

	utils.Check(container.Resolve(&controller), "Failed to create projectCategoryController instance...")

	router.Post("/project_categorys", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, project_category.CreateProjectCategoryValidator{})
	}, controller.Create)
	router.Get("/project_categorys", controller.GetAll)
	router.Get("/project_categorys/:id", controller.GetById)
	router.Delete("/project_categorys/:id", controller.Delete)
	router.Patch("/project_categorys/:id", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, project_category.UpdateProjectCategoryValidator{})
	}, controller.Update)
}
