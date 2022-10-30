package controllers

import "github.com/gofiber/fiber/v2"

type IAuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}
