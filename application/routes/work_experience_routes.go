package routes

import (
	schema_validation_middleware "github.com/gabemanfroi/schema-validation-middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/work_experience"
	"showcaseme/internal/utils"
)

func RegisterWorkExperienceRoutes(router fiber.Router) {
	var controller controllers.IWorkExperienceController

	utils.Check(container.Resolve(&controller), "Failed to create workExperienceController instance...")

	router.Post("/work_experiences", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, work_experience.CreateWorkExperienceValidator{})
	}, controller.Create)
	router.Get("/work_experiences", controller.GetAll)
	router.Get("/work_experiences/:id", controller.GetById)
	router.Delete("/work_experiences/:id", controller.Delete)
	router.Patch("/work_experiences/:id", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, work_experience.UpdateWorkExperienceValidator{})
	}, controller.Update)
}
