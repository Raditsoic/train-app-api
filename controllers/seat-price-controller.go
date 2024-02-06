package controllers

import (
	"github.com/Raditsoic/train-app-api/database"
	"github.com/Raditsoic/train-app-api/models"
	"github.com/gofiber/fiber/v2"
)

type SeatPriceController struct {
	BaseController
}

func (spc *SeatPriceController) GetAllbySeatID(c *fiber.Ctx) error {
    seatID := c.Params("seatID")
    var seatPrices []models.SeatPrice
    database.DB.Where("seat_id = ?", seatID).Find(&seatPrices)
    return c.JSON(seatPrices)
}

func (spc *SeatPriceController) GetAllbyJourneyID(c *fiber.Ctx) error {
    journeyID := c.Params("journeyID")
    var seatPrices []models.SeatPrice
    database.DB.Where("journey_id = ?", journeyID).Find(&seatPrices)
    return c.JSON(seatPrices)
}

func (spc *SeatPriceController) GetbySJID(c *fiber.Ctx) error {
    seatID := c.Query("seatID")
    journeyID := c.Query("journeyID")
    var seatPrice models.SeatPrice
    database.DB.Where("seat_id = ? AND journey_id = ?", seatID, journeyID).First(&seatPrice)
    return c.JSON(seatPrice)
}
