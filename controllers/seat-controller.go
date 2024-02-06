package controllers

import (
	"github.com/Raditsoic/train-app-api/database"
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
)

type SeatController struct {
	BaseController
}

func (sc *SeatController) GetAll(c *fiber.Ctx) error {
	var seats []models.Seat
	database.DB.Find(&seats)
	return c.JSON(seats)
}

func (sc *SeatController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var seat models.Seat
	if err := database.DB.First(&seat, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Seat not found")
	}
	return c.JSON(seat)
}

func (sc *SeatController) Create(c *fiber.Ctx) error {
	var seat models.Seat
	if err := c.BodyParser(&seat); err != nil {
		return err
	}
	database.DB.Create(&seat)
	return c.Status(fiber.StatusCreated).JSON(seat)
}

func (sc *SeatController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var seat models.Seat
	if err := database.DB.First(&seat, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Seat not found")
	}

	if err := c.BodyParser(&seat); err != nil {
		return err
	}

	database.DB.Save(&seat)
	return c.JSON(seat)
}

func (sc *SeatController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var seat models.Seat
	if err := database.DB.First(&seat, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Seat not found")
	}

	database.DB.Delete(&seat)
	return c.SendString("Seat deleted")
}
