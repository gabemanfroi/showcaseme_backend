package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/application/middlewares"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/skill"
	"showcaseme/internal/utils"
)

func RegisterSkillRoutes(router fiber.Router) {
	var controller controllers.ISkillController

	utils.Check(container.Resolve(&controller), "Failed to create skillController instance...")

	router.Post("/skills", func(c *fiber.Ctx) error { return middlewares.ValidateSchema(c, skill.CreateSkillValidator{}) }, controller.Create)
	router.Get("/skills", controller.GetAll)
	router.Get("/skills/:id", controller.GetById)
	router.Delete("/skills/:id", controller.Delete)
	router.Patch("/skills/:id", func(c *fiber.Ctx) error { return middlewares.ValidateSchema(c, skill.UpdateSkillValidator{}) }, controller.Update)
}
