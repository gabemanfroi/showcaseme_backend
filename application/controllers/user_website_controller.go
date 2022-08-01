package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/user_website"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type UserWebsiteController struct {
	service services.IUserWebsiteService
}

func CreateUserWebsiteController() *UserWebsiteController {
	return &UserWebsiteController{service: getUserWebsiteService()}
}

func getUserWebsiteService() services.IUserWebsiteService {
	var injector services.IUserWebsiteService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserWebsiteService instance ")
	return injector
}

func (controller UserWebsiteController) Create(c *fiber.Ctx) error {
	var dto user_website.CreateUserWebsiteDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdUserWebsite, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdUserWebsite)
}

func (controller UserWebsiteController) GetAll(c *fiber.Ctx) error {
	userWebsites, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&userWebsites), "failed to encode userWebsites")
	return c.Status(200).JSON(userWebsites)
}

func (controller UserWebsiteController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get user_website id")

	u, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&u), "failed to encode user_website")
	return c.Status(200).JSON(u)
}

func (controller UserWebsiteController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get user_website id")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("user_website deleted")
}

func (controller UserWebsiteController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get user_website id")
	var dto user_website.UpdateUserWebsiteDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedUserWebsite, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(updatedUserWebsite)
}
