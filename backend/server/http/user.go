package http

import (
	"github.com/DavidPsof/leetcode_problems/backend/pkg/api_models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := new(api_models.User)
		if err := ctx.BodyParser(user); err != nil {
			return ctx.Status(400).JSON(HandleError(err, 400, "CreateUser: Wrong input data format"))
		}

		return nil
	}
}
