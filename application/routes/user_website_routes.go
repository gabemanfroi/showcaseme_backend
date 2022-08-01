package routes

import (
	schema_validation_middleware "github.com/gabemanfroi/schema-validation-middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/user_website"
	"showcaseme/internal/utils"
)

func RegisterUserWebsiteRoutes(router fiber.Router) {
	var controller controllers.IUserWebsiteController

	utils.Check(container.Resolve(&controller), "Failed to create userWebsiteController instance...")

	router.Post("/user_websites", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, user_website.CreateUserWebsiteValidator{})
	}, controller.Create)
	router.Get("/user_websites", controller.GetAll)
	router.Get("/user_websites/:id", controller.GetById)
	router.Delete("/user_websites/:id", controller.Delete)
	router.Patch("/user_websites/:id", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, user_website.UpdateUserWebsiteValidator{})
	}, controller.Update)
}
