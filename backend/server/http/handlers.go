package http

import (
	"fmt"
	"github.com/DavidPsof/leetcode_problems/backend/pkg/api"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		if err := ctx.JSON(api.Response{
			Data: "pong",
		}); err != nil {
			return ctx.JSON(HandleError(err, 500, "Ping: Cant handle request"))
		}

		ctx.Status(http.StatusOK)

		return nil
	}
}

// HandleError - обработка ошибки в API методах
func HandleError(err error, errorType int, msg string) error {
	err = api.NewError(fmt.Sprintf("%s: %s", msg, err), errorType, "")
	log.Error(err)

	return api.NewError(msg, errorType, "")
}
