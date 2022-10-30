package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	UploadProfilePicture(c *fiber.Ctx) error
}
