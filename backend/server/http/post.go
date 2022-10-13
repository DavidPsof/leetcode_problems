package http

import (
	"github.com/DavidPsof/leetcode_problems/backend/pkg/api_models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// CreatePost upload md file of post and make record about it
func (h *Handler) CreatePost() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		uctx := ctx.UserContext()

		post := new(api_models.PostReq)
		post.Title = ctx.FormValue("title")

		file, err := ctx.FormFile("fileUpload")
		if err != nil {
			return ctx.Status(400).JSON(HandleError(err, 400, "CreatePost: Wrong input form file"))
		}

		if err = h.service.CreatePost(&uctx, post, file); err != nil {
			return ctx.Status(500).JSON(HandleError(err, 500, "CreatePost: server-side error"))
		}

		return nil
	}
}

// GetPost return one post
func (h *Handler) GetPost() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		uctx := ctx.UserContext()

		id, err := strconv.Atoi(ctx.Query("id"))

		post, err := h.service.GetPost(&uctx, id)
		if err != nil {
			return ctx.Status(500).JSON(HandleError(err, 500, "GetPost: server-side error"))
		}

		return ctx.JSON(post)
	}
}

// GetPosts return all posts
func (h *Handler) GetPosts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		posts, err := h.service.GetPosts()
		if err != nil {
			return ctx.Status(500).JSON(HandleError(err, 500, "GetPosts: server-side error"))
		}

		return ctx.JSON(posts)
	}
}
