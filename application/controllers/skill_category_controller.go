package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/skill_category"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type SkillCategoryController struct {
	service services.ISkillCategoryService
}

func CreateSkillCategoryController() *SkillCategoryController {
	return &SkillCategoryController{service: getSkillCategoryService()}
}

func getSkillCategoryService() services.ISkillCategoryService {
	var injector services.ISkillCategoryService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}

func (controller SkillCategoryController) Create(c *fiber.Ctx) error {
	var dto skill_category.CreateSkillCategoryDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdSkillCategory, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdSkillCategory)

}

func (controller SkillCategoryController) GetAll(c *fiber.Ctx) error {
	users, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&users), "failed to encode users")
	return c.Status(200).JSON(users)
}

func (controller SkillCategoryController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")

	user, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&user), "failed to encode user")
	return c.Status(200).JSON(user)
}

func (controller SkillCategoryController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("skill_category deleted")
}

func (controller SkillCategoryController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")

	var dto skill_category.UpdateSkillCategoryDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedUser, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(updatedUser)
}
