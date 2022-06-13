package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
)

type UserController struct {
	service services.UserServiceInterface
}

func CreateUserController() *UserController { return &UserController{service: getService()} }

func (u UserController) Create(c *fiber.Ctx) error {
	var dto user.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	return c.Status(200).JSON(u.service.Create(&dto))
}

func (u UserController) GetAll(c *fiber.Ctx) error {
	users, err := u.service.GetAll()
	if err != nil {
		return err
	}
	json.NewEncoder(c).Encode(&users)
	return c.Status(200).JSON(users)
}

func (u UserController) GetById(c *fiber.Ctx) error {
	user, err := u.service.GetById(c.Params("id"))
	if err != nil {
		return err
	}
	json.NewEncoder(c).Encode(&user)
	return c.Status(200).JSON(user)
}

func (u UserController) Delete(c *fiber.Ctx) error {
	err := u.service.Delete(c.Params("id"))
	if err != nil {
		return err
	}

	return c.Status(200).JSON("user deleted")
}

func (u UserController) Update(c *fiber.Ctx) error {
	var dto user.UpdateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedUser, err := u.service.Update(c.Params("id"), &dto)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(updatedUser)
}

func getService() services.UserServiceInterface {
	var injector services.UserServiceInterface
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}
