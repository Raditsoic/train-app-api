package controllers

import (
	"github.com/Raditsoic/train-app-api/database"
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
)

type CarriageController struct {
	BaseController
}

func (cc *CarriageController) GetAll(c *fiber.Ctx) error {
	var carriages []models.Carriage
	database.DB.Find(&carriages)
	return c.JSON(carriages)
}

func (cc *CarriageController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var carriage models.Carriage
	if err := database.DB.First(&carriage, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Carriage not found")
	}
	return c.JSON(carriage)
}

func (cc *CarriageController) Create(c *fiber.Ctx) error {
	var carriage models.Carriage
	if err := c.BodyParser(&carriage); err != nil {
		return err
	}
	database.DB.Create(&carriage)
	return c.Status(fiber.StatusCreated).JSON(carriage)
}

func (cc *CarriageController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var carriage models.Carriage
	if err := database.DB.First(&carriage, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Carriage not found")
	}

	if err := c.BodyParser(&carriage); err != nil {
		return err
	}

	database.DB.Save(&carriage)
	return c.JSON(carriage)
}

func (cc *CarriageController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var carriage models.Carriage
	if err := database.DB.First(&carriage, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Carriage not found")
	}

	database.DB.Delete(&carriage)
	return c.SendString("Carriage deleted")
}
