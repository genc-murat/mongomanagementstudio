package router

import (
	"mongomanagementstudio/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func HTMLRouter(handler *handler.HTMLHandler, app *fiber.App) {
	app.Get("/", handler.Index)
}

func CollectionStatsRouter(handler *handler.CollectionStatsHandler, app *fiber.App) {
	collectionStatsRouter := app.Group("/collectionstats")
	collectionStatsRouter.Get("/:collection", handler.CollectionStats)
}
