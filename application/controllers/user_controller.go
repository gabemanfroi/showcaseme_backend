package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/user"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type UserController struct {
	service services.IUserService
}

func CreateUserController() *UserController { return &UserController{service: getUserService()} }

func (controller UserController) Create(c *fiber.Ctx) error {
	var dto user.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdUser, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdUser)
}

func (controller UserController) GetAll(c *fiber.Ctx) error {
	users, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&users), "failed to encode users")
	return c.Status(200).JSON(users)
}

func (controller UserController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")
	u, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&u), "failed to encode user")
	return c.Status(200).JSON(u)
}

func (controller UserController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")

	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("user deleted")
}

func (controller UserController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")
	var dto user.UpdateUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedUser, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(updatedUser)
}

func (controller UserController) UploadProfilePicture(c *fiber.Ctx) error {
	username := c.FormValue("username")
	profilePicture, err := c.FormFile("profilePicture")
	if err != nil {
		return c.Status(400).JSON(err)
	}
	response, err := controller.service.UploadProfilePicture(username, profilePicture)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(response)
}

func getUserService() services.IUserService {
	var injector services.IUserService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}
