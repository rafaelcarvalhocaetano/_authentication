package main

import (
	db "go_auth/database"

	rt "go_auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.ConnectionPostgres()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//routers
	rt.Setup(app)

	app.Listen(":8000")
}
