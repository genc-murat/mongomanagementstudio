package handler

import (
	"mongomanagementstudio/internal/store"

	"github.com/gofiber/fiber/v2"
)

type HTMLHandler struct {
	repo store.Repository
}

func NewHTMLHandler(repo store.Repository) *HTMLHandler {
	return &HTMLHandler{repo}
}

func (h *HTMLHandler) Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Mongo Management Studio",
	})
}
