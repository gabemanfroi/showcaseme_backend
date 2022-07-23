package routes

import (
	schema_validation_middleware "github.com/gabemanfroi/schema-validation-middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/infra/validators/skill_category"
	"showcaseme/internal/utils"
)

func RegisterSkillCategoryCategoryRoutes(router fiber.Router) {
	var controller controllers.ISkillCategoryController

	utils.Check(container.Resolve(&controller), "Failed to create skill_categoryController instance")

	router.Post("/skill_categories", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, skill_category.CreateSkillCategoryValidator{})
	}, controller.Create)
	router.Get("/skill_categories", controller.GetAll)
	router.Get("/skill_categories/:id", controller.GetById)
	router.Delete("/skill_categories/:id", controller.Delete)
	router.Patch("/skill_categories/:id", func(c *fiber.Ctx) error {
		return schema_validation_middleware.ValidateSchema(c, skill_category.UpdateSkillCategoryValidator{})
	}, controller.Update)
}
