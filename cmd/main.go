package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)

	}

	db := client.Database("mongomanagementstudio")
	//collection := db.Collection("names")
	result := db.RunCommand(context.Background(), bson.M{"collStats": "names"})

	var document bson.M
	err = result.Decode(&document)

	if err != nil {
		panic(err)
	}

	fmt.Printf("REUSABLE: %v Bytes\n", document["wiredTiger"])
	fmt.Printf("Collection size: %v Bytes\n", document["size"])
	fmt.Printf("Average object size: %v Bytes\n", document["avgObjSize"])
	fmt.Printf("Storage size: %v Bytes\n", document["storageSize"])
	fmt.Printf("Total index size: %v Bytes\n", document["totalIndexSize"])

}
