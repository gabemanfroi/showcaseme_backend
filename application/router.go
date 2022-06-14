package application

import (
	"github.com/gofiber/fiber/v2"
	"showcaseme/application/routes"
)

func RegisterRoutes(router fiber.Router) {
	routes.RegisterUserRoutes(router)
	routes.RegisterSkillRoutes(router)
}
