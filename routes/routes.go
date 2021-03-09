package routes

import (
	"demo/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/logout", controller.Logout)

	app.Get("/api/user", controller.User)
}
