package controllers

import "github.com/gofiber/fiber/v2"

type Controller interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Patch(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type BaseController struct{}

func (bc *BaseController) GetAll(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Get method not implemented")
}

func (bc *BaseController) Post(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Create method not implemented")
}

func (bc *BaseController) Delete(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Delete method not implemented")
}

func (bc *BaseController) Patch(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Update method not implemented")
}

func (bc *BaseController) GetByID(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotImplemented, "Get by ID method not implemented")
}
