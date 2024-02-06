package controllers

import (
	"github.com/Raditsoic/train-app-api/database"
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
)

type BookingController struct {
	BaseController
}

func (bc *BookingController) GetAll(c *fiber.Ctx) error {
	var bookings []models.Booking
	database.DB.Find(&bookings)
	return c.JSON(bookings)
}

func (bc *BookingController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var booking models.Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Booking not found")
	}
	return c.JSON(booking)
}

func (bc *BookingController) Create(c *fiber.Ctx) error {
	var booking models.Booking
	if err := c.BodyParser(&booking); err != nil {
		return err
	}
	database.DB.Create(&booking)
	return c.Status(fiber.StatusCreated).JSON(booking)
}

func (bc *BookingController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var booking models.Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Booking not found")
	}

	if err := c.BodyParser(&booking); err != nil {
		return err
	}

	database.DB.Save(&booking)
	return c.JSON(booking)
}

func (bc *BookingController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var booking models.Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Booking not found")
	}

	database.DB.Delete(&booking)
	return c.SendString("Booking deleted")
}
