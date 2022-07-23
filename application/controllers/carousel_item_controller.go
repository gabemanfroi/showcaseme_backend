package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/carousel_item"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type CarouselItemController struct {
	service services.ICarouselItemService
}

func CreateCarouselItemController() *CarouselItemController {
	return &CarouselItemController{service: getCarouselItemService()}
}

func getCarouselItemService() services.ICarouselItemService {
	var injector services.ICarouselItemService
	utils.Check(container.Resolve(&injector), "Error while retrieving UserService instance ")
	return injector
}

func (controller CarouselItemController) Create(c *fiber.Ctx) error {
	var dto carousel_item.CreateCarouselItemDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdCarouselItem, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdCarouselItem)

}

func (controller CarouselItemController) GetAll(c *fiber.Ctx) error {
	users, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&users), "failed to encode users")
	return c.Status(200).JSON(users)
}

func (controller CarouselItemController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")

	user, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&user), "failed to encode user")
	return c.Status(200).JSON(user)
}

func (controller CarouselItemController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("carousel_item deleted")
}

func (controller CarouselItemController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get idParams")

	var dto carousel_item.UpdateCarouselItemDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedUser, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(updatedUser)
}
