package http

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	app.Get("/ping", Ping())
}
