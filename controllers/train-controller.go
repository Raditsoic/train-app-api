package controllers

import (
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/Raditsoic/train-app-api/database"
)

type TrainController struct {
	BaseController
}

func (tc *TrainController) GetAll(c *fiber.Ctx) error {
    var trains []models.Train
    database.DB.Find(&trains)
    return c.JSON(trains)
}

func (tc *TrainController) GetByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var train models.Train
    if err := database.DB.First(&train, id).Error; err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Train not found")
    }
    return c.JSON(train)
}

func (tc *TrainController) Create(c *fiber.Ctx) error {
    var train models.Train
    if err := c.BodyParser(&train); err != nil {
        return err
    }
    database.DB.Create(&train)
    return c.Status(fiber.StatusCreated).JSON(train)
}

func (tc *TrainController) Update(c *fiber.Ctx) error {
    id := c.Params("id")
    var train models.Train
    if err := database.DB.First(&train, id).Error; err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Train not found")
    }

    if err := c.BodyParser(&train); err != nil {
        return err
    }

    database.DB.Save(&train)
    return c.JSON(train)
}

func (tc *TrainController) Delete(c *fiber.Ctx) error {
    id := c.Params("id")
    var train models.Train
    if err := database.DB.First(&train, id).Error; err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Train not found")
    }

    database.DB.Delete(&train)
    return c.SendString("Train deleted")
}