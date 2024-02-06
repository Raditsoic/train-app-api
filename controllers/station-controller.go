package controllers

import (
	"github.com/Raditsoic/train-app-api/database"
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
)

type StationController struct {
	BaseController
}

func (sc *StationController) GetAll(c *fiber.Ctx) error {
	var stations []models.Station
	database.DB.Find(&stations)
	return c.JSON(stations)
}

func (sc *StationController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var station models.Station
	if err := database.DB.First(&station, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Station not found")
	}
	return c.JSON(station)
}

func (sc *StationController) Create(c *fiber.Ctx) error {
	var station models.Station
	if err := c.BodyParser(&station); err != nil {
		return err
	}
	database.DB.Create(&station)
	return c.Status(fiber.StatusCreated).JSON(station)
}

func (sc *StationController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var station models.Station
	if err := database.DB.First(&station, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Station not found")
	}

	if err := c.BodyParser(&station); err != nil {
		return err
	}

	database.DB.Save(&station)
	return c.JSON(station)
}

func (sc *StationController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var station models.Station
	if err := database.DB.First(&station, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Station not found")
	}

	database.DB.Delete(&station)
	return c.SendString("Station deleted")
}
