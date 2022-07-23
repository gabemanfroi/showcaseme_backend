package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
)

type ResumeController struct {
	service services.IResumeService
}

func CreateResumeController() *ResumeController {
	return &ResumeController{service: getResumeService()}
}

func getResumeService() services.IResumeService {
	var injector services.IResumeService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}

func (controller ResumeController) GetByUsername(c *fiber.Ctx) error {
	resume, err := controller.service.GetByUsername(c.Params("username"))

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	utils.Check(json.NewEncoder(c).Encode(&resume), "failed to encode resume")

	return c.Status(200).JSON(resume)
}
