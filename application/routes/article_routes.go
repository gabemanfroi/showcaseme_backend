package routes

import (
	schema_validation_middleware "github.com/gabemanfroi/schema-validation-middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/article"
	"showcaseme/internal/utils"
)

func RegisterArticleRoutes(router fiber.Router) {
	var controller controllers.IArticleController

	utils.Check(container.Resolve(&controller), "Failed to create articleController instance...")

	router.Post("/articles", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, article.CreateArticleValidator{})
	}, controller.Create)
	router.Get("/articles", controller.GetAll)
	router.Get("/articles/:id", controller.GetById)
	router.Delete("/articles/:id", controller.Delete)
	router.Patch("/articles/:id", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, article.UpdateArticleValidator{})
	}, controller.Update)
}
