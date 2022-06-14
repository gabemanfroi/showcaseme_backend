package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
)

type SkillController struct {
	service services.ISkillService
}

func (controller SkillController) Create(c *fiber.Ctx) error {
	var dto skill.CreateSkillDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	return c.Status(200).JSON(controller.service.Create(&dto))

}

func (controller SkillController) GetAll(c *fiber.Ctx) error {
	users, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&users), "failed to encode users")
	return c.Status(200).JSON(users)
}

func (controller SkillController) GetById(c *fiber.Ctx) error {
	user, err := controller.service.GetById(c.Params("id"))
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&user), "failed to encode user")
	return c.Status(200).JSON(user)
}

func (controller SkillController) Delete(c *fiber.Ctx) error {
	err := controller.service.Delete(c.Params("id"))
	if err != nil {
		return err
	}

	return c.Status(200).JSON("user deleted")
}

func (controller SkillController) Update(c *fiber.Ctx) error {
	var dto skill.UpdateSkillDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedUser, err := controller.service.Update(c.Params("id"), &dto)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(updatedUser)
}

func CreateSkillController() *SkillController {
	return &SkillController{service: getSkillService()}
}

func getSkillService() services.ISkillService {
	var injector services.ISkillService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}
