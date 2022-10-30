package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/internal/utils"
)

func RegisterAuthRoutes(router fiber.Router) {
	var controller controllers.IAuthController

	utils.Check(container.Resolve(&controller), "Failed to create carousel_itemController instance")
	router.Post("/auth/login", controller.Login)
	router.Post("/auth/register", controller.Register)
}
