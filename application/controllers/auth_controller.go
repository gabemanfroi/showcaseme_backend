package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/auth"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
)

type AuthController struct {
	service services.IAuthService
}

func CreateAuthController() *AuthController {
	return &AuthController{service: getAuthService()}
}

func (controller AuthController) Login(c *fiber.Ctx) error {
	var dto auth.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	if err := controller.service.Login(&dto); err != nil {
		return c.Status(401).JSON(err.Error())
	}
	return nil
}

func (controller AuthController) Register(c *fiber.Ctx) error {
	var dto auth.RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	if _, err := controller.service.Register(&dto); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(201).JSON(nil)
}

func getAuthService() services.IAuthService {
	var injector services.IAuthService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}
