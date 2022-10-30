package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/project"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type ProjectController struct {
	service services.IProjectService
}

func CreateProjectController() *ProjectController {
	return &ProjectController{service: getProjectService()}
}

func getProjectService() services.IProjectService {
	var injector services.IProjectService
	utils.Check(container.Resolve(&injector), "Error while retrieving ProjectService instance ")
	return injector
}

func (controller ProjectController) Create(c *fiber.Ctx) error {
	backgroundImage, err := c.FormFile("backgroundImage")
	if err != nil {
		return c.Status(400).JSON(err)
	}

	var dto project.CreateProjectDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	dto.BackgroundImage = backgroundImage

	fmt.Println(dto)

	createdProject, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdProject)
}

func (controller ProjectController) GetAll(c *fiber.Ctx) error {
	projects, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&projects), "failed to encode projects")
	return c.Status(200).JSON(projects)
}

func (controller ProjectController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get project id")

	p, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&p), "failed to encode project")
	return c.Status(200).JSON(p)
}

func (controller ProjectController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get project id")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("project deleted")
}

func (controller ProjectController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get project id")
	var dto project.UpdateProjectDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedProject, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(updatedProject)
}
