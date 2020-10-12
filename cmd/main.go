package main

import (
	"context"

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
	// ctx := context.Background()
	// mongoStore, err := driver.NewMongoStore(ctx, "mongodb://localhost:27017/mongomanagementstudio", "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// collectionStatsRepo, err := repository.NewCommandRepository(mongoStore, "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// collectionStatsHandler := handler.NewCollectionStatsHandler(collectionStatsRepo)

	// app := fiber.New()
	// router.CollectionStatsRouter(collectionStatsHandler, app)
	// app.Listen(":5353")
	// return nil
}
