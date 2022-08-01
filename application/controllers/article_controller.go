package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
	"showcaseme/domain/DTO/article"
	"showcaseme/domain/interfaces/services"
	"showcaseme/internal/utils"
	"strconv"
)

type ArticleController struct {
	service services.IArticleService
}

func CreateArticleController() *ArticleController {
	return &ArticleController{service: getArticleService()}
}

func getArticleService() services.IArticleService {
	var injector services.IArticleService
	utils.Check(container.Resolve(&injector), "Error while retrieving ArticleService instance ")
	return injector
}

func (controller ArticleController) Create(c *fiber.Ctx) error {
	var dto article.CreateArticleDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	createdArticle, err := controller.service.Create(&dto)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(createdArticle)
}

func (controller ArticleController) GetAll(c *fiber.Ctx) error {
	articles, err := controller.service.GetAll()
	if err != nil {
		return err
	}
	utils.Check(json.NewEncoder(c).Encode(&articles), "failed to encode articles")
	return c.Status(200).JSON(articles)
}

func (controller ArticleController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get article id")

	a, err := controller.service.GetById(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	utils.Check(json.NewEncoder(c).Encode(&a), "failed to encode article")
	return c.Status(200).JSON(a)
}

func (controller ArticleController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get article id")
	err = controller.service.Delete(uint(id))
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("article deleted")
}

func (controller ArticleController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	utils.Check(err, "failed to get article id")
	var dto article.UpdateArticleDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	updatedArticle, err := controller.service.Update(uint(id), &dto)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(updatedArticle)
}
