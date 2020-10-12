package router

import (
	"mongomanagementstudio/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func HTMLRouter(handler *handler.HTMLHandler, app *fiber.App) {
	app.Get("/", handler.Index)
}

func CollectionStatsRouter(handler *handler.CollectionStatsHandler, app *fiber.App) {
	collectionStatsRouter := app.Group("/collection")
	//collection/stats/denemeCollection
	collectionStatsRouter.Get("/stats/:name", handler.CollectionStats)
	collectionStatsRouter.Get("/list/:server/:port/:db", handler.CollectionNames)
}
