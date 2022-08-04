package controllers

import (
	"encoding/json"
	"showcaseme/domain/DTO/work_experience"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"strconv"
)

type WorkExperienceController struct {
	service services.IWorkExperienceService
}

func CreateWorkExperienceController() *WorkExperienceController {
    return &WorkExperienceController{service: getWorkExperienceService()}
}

func getWorkExperienceService() services.IWorkExperienceService {
	var injector services.IWorkExperienceService
	utils.Check(container.Resolve(&injector), "Error while retrieving WorkExperienceService instance ")
	return injector
}

func (controller WorkExperienceController) Create(c *fiber.Ctx) error {
	var dto work_experience.CreateWorkExperienceDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdWorkExperience, err := controller.service.Create(&dto)

	if err != nil {
    		return c.Status(400).JSON(err.Error())
    }

	return c.Status(200).JSON(createdWorkExperience)
}

func (controller WorkExperienceController) GetAll(c *fiber.Ctx) error {
	workExperiences, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&workExperiences), "failed to encode workExperiences")
	return c.Status(200).JSON(workExperiences)
}

func (controller WorkExperienceController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get work_experience id")

	w, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&w), "failed to encode work_experience")
	return c.Status(200).JSON(w)
}

func (controller WorkExperienceController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get work_experience id")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(204).JSON("user deleted")
}

func (controller WorkExperienceController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get work_experience id")
	var dto work_experience.UpdateWorkExperienceDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedWorkExperience, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(updatedWorkExperience)
}
