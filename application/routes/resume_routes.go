package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/controllers"
	"showcaseme/internal/utils"
)

func RegisterResumeRoutes(router fiber.Router) {
	var controller controllers.IResumeController

	utils.Check(container.Resolve(&controller), "Failed to create resumeController instance")
	router.Get("/resumes/:username", controller.GetByUsername)
}
