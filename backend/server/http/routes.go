package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Get("/ping", h.Ping())

	users := app.Group("/user")
	users.Post("register", h.CreateUser())

	posts := app.Group("/posts")
	posts.Get("", h.GetPost())
	posts.Get("/all", h.GetPosts())
	posts.Post("", h.CreatePost())

}
