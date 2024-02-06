package controller

import (
	"github.com/Raditsoic/train-app-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	SetupBookingRoutes(app, &controllers.BookingController{})
	SetupCarriageRoutes(app, &controllers.CarriageController{})
	SetupJourneyRoutes(app, &controllers.JourneyController{})
	SetupSeatPriceRoutes(app, &controllers.SeatPriceController{})
	SetupSeatRoutes(app, &controllers.SeatController{})
	SetupStationRoutes(app, &controllers.StationController{})
	SetupUserRoutes(app, &controllers.UserController{})
}

func SetupTrainRoutes(app *fiber.App, tc *controllers.TrainController) {
	train := app.Group("/trains")
	train.Get("/", tc.GetAll)
	train.Post("/", tc.Create)
	train.Delete("/:id", tc.Delete)
	train.Patch("/:id", tc.Update)
	train.Get("/:id", tc.GetByID)
}

func SetupBookingRoutes(app *fiber.App, bc *controllers.BookingController) {
	booking := app.Group("/bookings")
	booking.Get("/", bc.GetAll)
	booking.Post("/", bc.Create)
	booking.Get("/:id", bc.GetByID)
	booking.Patch("/:id", bc.Patch)
	booking.Delete("/:id", bc.Delete)
}

func SetupCarriageRoutes(app *fiber.App, cc *controllers.CarriageController) {
	carriage := app.Group("/carriages")
	carriage.Get("/", cc.GetAll)
	carriage.Post("/", cc.Create)
	carriage.Get("/:id", cc.GetByID)
	carriage.Patch("/:id", cc.Patch)
	carriage.Delete("/:id", cc.Delete)
}

func SetupJourneyRoutes(app *fiber.App, jc *controllers.JourneyController) {
	journey := app.Group("/journeys")
	journey.Get("/", jc.GetAll)
	journey.Post("/", jc.Create)
	journey.Get("/:id", jc.GetByID)
	journey.Patch("/:id", jc.Patch)
	journey.Delete("/:id", jc.Delete)
}

func SetupSeatRoutes(app *fiber.App, sc *controllers.SeatController) {
	seat := app.Group("/seats")
	seat.Get("/", sc.GetAll)
	seat.Post("/", sc.Create)
	seat.Get("/:id", sc.GetByID)
	seat.Patch("/:id", sc.Patch)
	seat.Delete("/:id", sc.Delete)
}

func SetupSeatPriceRoutes(app *fiber.App, spc *controllers.SeatPriceController) {
	seat_price := app.Group("/seat-prices")
	seat_price.Post("/", spc.Post)
	seat_price.Get("/journeys/:journeyID", spc.GetAllbyJourneyID)
	seat_price.Get("/seats/:seatID", spc.GetAllbySeatID)
	seat_price.Get("/:seatID-:journeyID", spc.GetbySJID)
}

func SetupStationRoutes(app *fiber.App, sc *controllers.StationController) {
	station := app.Group("/stations")
	station.Get("/", sc.GetAll)
	station.Post("/", sc.Create)
	station.Get("/:id", sc.GetByID)
	station.Patch("/:id", sc.Patch)
	station.Delete("/:id", sc.Delete)
}

func SetupUserRoutes(app *fiber.App, uc *controllers.UserController) {
	user := app.Group("/users")
	user.Get("/", uc.GetAll)
	user.Post("/", uc.Post)
	user.Get("/:id", uc.GetByID)
	user.Patch("/:id", uc.Patch)
	user.Delete("/:id", uc.Delete)
}
