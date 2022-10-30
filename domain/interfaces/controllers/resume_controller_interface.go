package controllers

import "github.com/gofiber/fiber/v2"

type IResumeController interface {
	GetByUsername(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
