package routes

import (
	controller "go_auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controller.HiApi)
	app.Get("api/user", controller.User)

	app.Post("api/register", controller.Register)
	app.Post("api/login", controller.Login)
	app.Post("api/logout", controller.Logout)

}
