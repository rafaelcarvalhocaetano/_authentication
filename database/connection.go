package database

import (
	"go_auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionPostgres() {
	dsn := "host=postgres user=root password=postgres dbname=go_auth port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
