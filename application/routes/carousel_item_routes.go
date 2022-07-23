package routes

import (
	schema_validation_middleware "github.com/gabemanfroi/schema-validation-middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/carousel_item"
	"showcaseme/internal/utils"
)

func RegisterCarouselItemCategoryRoutes(router fiber.Router) {
	var controller controllers.ICarouselItemController

	utils.Check(container.Resolve(&controller), "Failed to create carousel_itemController instance")

	router.Post("/carousel_items", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, carousel_item.CreateCarouselItemValidator{})
	}, controller.Create)
	router.Get("/carousel_items", controller.GetAll)
	router.Get("/carousel_items/:id", controller.GetById)
	router.Delete("/carousel_items/:id", controller.Delete)
	router.Patch("/carousel_items/:id", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, carousel_item.UpdateCarouselItemValidator{})
	}, controller.Update)
}
