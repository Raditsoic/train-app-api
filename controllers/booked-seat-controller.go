package controllers

import (
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/Raditsoic/train-app-api/database"
)

type BookedSeatController struct {
	BaseController
}

func (tc *BookedSeatController) GetAllbySeatID(c *fiber.Ctx) error {
    var seat_price []models.BookedSeat
    database.DB.Find(&seat_price)/*Not implemented*/
    return c.JSON(seat_price)
}

func (spc *BookedSeatController) GetAllbyBookingID(c *fiber.Ctx) error {
	var seat_price []models.BookedSeat
	database.DB.Find(&seat_price)
	return c.JSON(seat_price)
}

func (spc *BookedSeatController) GetAll(c *fiber.Ctx) error {
	return c.JSON("Method not yet Implemented")
}

func (spc *BookedSeatController) Create(c *fiber.Ctx) error {
	var booking models.Booking
	if err := c.BodyParser(&booking); err != nil {
		return err
	}
	database.DB.Create(&booking)
	return c.Status(fiber.StatusCreated).JSON(booking)
}