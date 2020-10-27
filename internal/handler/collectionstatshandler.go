package handler

import (
	"context"
	"fmt"
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
	collection := c.Params("name")
	result := h.repo.RunCommand(context.Background(), bson.M{"collStats": collection})

	data, err := bson.Marshal(result)
	if err != nil {

	}

	var collectionStats models.CollectionStats
	err = bson.Unmarshal(data, &collectionStats)

	c.JSON(collectionStats)
	return nil
}

func (h *CollectionStatsHandler) CollectionNames(c *fiber.Ctx) error {
	server := c.Params("server")
	port := c.Params("port")
	db := c.Params("db")
	h.repo.SetConnectionString("mongodb://" + server + ":" + port + "/" + db)

	names, err := h.repo.CollectionNames(context.Background())
	c.JSON(names)
	return err
}

func (h *CollectionStatsHandler) Compact(c *fiber.Ctx) error {
	collection := c.Params("name")
	result := h.repo.RunCommand(context.Background(), bson.M{"compact": collection})
	fmt.Println(result["bytesFreed"])
	if result != nil {

	}
	//c.JSON(fiber.Map{"bytesFreed": result["bytesFreed"], "ok": result["ok"]})
	c.JSON(result)
	return nil
}
