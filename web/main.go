package main

import (
	"context"
	"fmt"
	"mongomanagementstudio/internal/handler"
	"mongomanagementstudio/internal/repository"
	"mongomanagementstudio/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	runCollectionStats()
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)

	// }

	// db := client.Database("mongomanagementstudio")
	// //collection := db.Collection("names")
	// result := db.RunCommand(context.Background(), bson.M{"collStats": "names"})

	// //var document bson.M
	// var document models.CollectionStats
	// err = result.Decode(&document)

	// if err != nil {
	// 	panic(err)
	// }

	// // fmt.Printf("REUSABLE: %v Bytes\n", document["wiredTiger"])
	// fmt.Printf("Collection size: %v Bytes\n", document.WiredTiger.BlockManager.FileBytesAvailableForReUse)
	// // fmt.Printf("Average object size: %v Bytes\n", document["avgObjSize"])
	// // fmt.Printf("Storage size: %v Bytes\n", document["storageSize"])
	// // fmt.Printf("Total index size: %v Bytes\n", document["totalIndexSize"])

}

func runCollectionStats() error {
	//ctx := context.Background()
	// mongoStore, err := driver.NewMongoStore(ctx, "mongodb://localhost:27017/mongomanagementstudio", "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	collectionStatsRepo, err := repository.NewCommandRepository(nil, "")
	if err != nil {
		fmt.Println(err)
	}

	collectionStatsHandler := handler.NewCollectionStatsHandler(collectionStatsRepo)
	htmlHandler := handler.NewHTMLHandler(collectionStatsRepo)

	engine := django.New("web/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/public/css", "./web/static/css")
	app.Static("/public/img", "./web/static/img")
	app.Static("/public/js", "./web/static/js")
	router.CollectionStatsRouter(collectionStatsHandler, app)
	router.HTMLRouter(htmlHandler, app)
	app.Listen(":5353")
	return nil
}
