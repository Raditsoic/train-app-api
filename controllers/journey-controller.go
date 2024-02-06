package controllers

import (
	"github.com/Raditsoic/train-app-api/database"
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
)

type JourneyController struct {
	BaseController
}

func (jc *JourneyController) GetAll(c *fiber.Ctx) error {
	var journeys []models.Journey
	database.DB.Find(&journeys)
	return c.JSON(journeys)
}

func (jc *JourneyController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var journey models.Journey
	if err := database.DB.First(&journey, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Journey not found")
	}
	return c.JSON(journey)
}

func (jc *JourneyController) Create(c *fiber.Ctx) error {
	var journey models.Journey
	if err := c.BodyParser(&journey); err != nil {
		return err
	}
	database.DB.Create(&journey)
	return c.Status(fiber.StatusCreated).JSON(journey)
}

func (jc *JourneyController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var journey models.Journey
	if err := database.DB.First(&journey, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Journey not found")
	}

	if err := c.BodyParser(&journey); err != nil {
		return err
	}

	database.DB.Save(&journey)
	return c.JSON(journey)
}

func (jc *JourneyController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var journey models.Journey
	if err := database.DB.First(&journey, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Journey not found")
	}

	database.DB.Delete(&journey)
	return c.SendString("Journey deleted")
}
