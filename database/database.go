package database

import (
	"fmt"

	"github.com/Raditsoic/train-app-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	host := "localhost"
	port := "5432"
	user := "raditsoic"
	password := "123"
	dbname := "movie_db"

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = database
	return nil
}

func Migrate() error {
	if err := DB.AutoMigrate(&models.Train{}, &models.Station{}, &models.Carriage{}, 
		&models.Seat{}, &models.SeatPrice{}, &models.Booking{}, 
		&models.Journey{}, &models.User{}); err != nil {
		return err
	}

	return nil
}
