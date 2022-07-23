package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/skill"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type SkillController struct {
	service services.ISkillService
}

func CreateSkillController() *SkillController {
	return &SkillController{service: getSkillService()}
}

func getSkillService() services.ISkillService {
	var injector services.ISkillService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}

func (controller SkillController) Create(c *fiber.Ctx) error {
	var dto skill.CreateSkillDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdSkill, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdSkill)

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
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")

	user, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&user), "failed to encode user")
	return c.Status(200).JSON(user)
}

func (controller SkillController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("user deleted")
}

func (controller SkillController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")

	var dto skill.UpdateSkillDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	updatedUser, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(updatedUser)
}
