package database

import (
	"go-belajar/config"
	"go-belajar/models"
)

func Migrate() {
	config.DB.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Booking{},
	)
}
