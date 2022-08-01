package routes

import (
	schema_validation_middleware "github.com/gabemanfroi/schema-validation-middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/project"
	"showcaseme/internal/utils"
)

func RegisterProjectRoutes(router fiber.Router) {
	var controller controllers.IProjectController

	utils.Check(container.Resolve(&controller), "Failed to create projectController instance...")

	router.Post("/projects", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, project.CreateProjectValidator{})
	}, controller.Create)
	router.Get("/projects", controller.GetAll)
	router.Get("/projects/:id", controller.GetById)
	router.Delete("/projects/:id", controller.Delete)
	router.Patch("/projects/:id", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, project.UpdateProjectValidator{})
	}, controller.Update)
}
