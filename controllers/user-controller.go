package controllers

import (
	"github.com/Raditsoic/train-app-api/database"
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	BaseController
}

func (uc *UserController) GetAll(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func (uc *UserController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	return c.JSON(user)
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Save(&user)
	return c.JSON(user)
}

func (uc *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	database.DB.Delete(&user)
	return c.SendString("user deleted")
}
