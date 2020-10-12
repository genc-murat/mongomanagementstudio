package handler

import (
	"context"
	"mongomanagementstudio/internal/models"
	"mongomanagementstudio/internal/store"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type CollectionStatsHandler struct {
	repo store.Repository
}

func NewCollectionStatsHandler(repo store.Repository) *CollectionStatsHandler {
	return &CollectionStatsHandler{repo}
}

func (h *CollectionStatsHandler) CollectionStats(c *fiber.Ctx) error {
	//collection := c.Params("collection")
	result := h.repo.RunCommand(context.TODO(), bson.M{"collStats": "names"})
	data, err := bson.Marshal(result)
	if err != nil {

	}
	var collectionStats models.CollectionStats
	err = bson.Unmarshal(data, &collectionStats)
	c.JSON(collectionStats)
	return nil
}
