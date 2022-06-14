package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/application/middlewares"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/user"
	"showcaseme/internal/utils"
)

func RegisterUserRoutes(router fiber.Router) {
	var controller controllers.IUserController

	utils.Check(container.Resolve(&controller), "Failed to create userController instance...")

	router.Post("/users", func(c *fiber.Ctx) error { return middlewares.ValidateSchema(c, user.CreateUserValidator{}) }, controller.Create)
	router.Get("/users", controller.GetAll)
	router.Get("/users/:id", controller.GetById)
	router.Delete("/users/:id", controller.Delete)
	router.Patch("/users/:id", func(c *fiber.Ctx) error { return middlewares.ValidateSchema(c, user.UpdateUserValidator{}) }, controller.Update)
}
