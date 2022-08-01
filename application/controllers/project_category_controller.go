package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/project_category"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type ProjectCategoryController struct {
	service services.IProjectCategoryService
}

func CreateProjectCategoryController() *ProjectCategoryController {
	return &ProjectCategoryController{service: getProjectCategoryService()}
}

func getProjectCategoryService() services.IProjectCategoryService {
	var injector services.IProjectCategoryService
	utils.Check(container.Resolve(&injector), "Error while retrieving ProjectCategoryService instance ")
	return injector
}

func (controller ProjectCategoryController) Create(c *fiber.Ctx) error {
	var dto project_category.CreateProjectCategoryDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdProjectCategory, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdProjectCategory)
}

func (controller ProjectCategoryController) GetAll(c *fiber.Ctx) error {
	projectCategorys, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&projectCategorys), "failed to encode projectCategories")
	return c.Status(200).JSON(projectCategorys)
}

func (controller ProjectCategoryController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get project_category id")

	p, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&p), "failed to encode project_category")
	return c.Status(200).JSON(p)
}

func (controller ProjectCategoryController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get project_category id")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("project_category deleted")
}

func (controller ProjectCategoryController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get project_category id")
	var dto project_category.UpdateProjectCategoryDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedProjectCategory, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(updatedProjectCategory)
}
