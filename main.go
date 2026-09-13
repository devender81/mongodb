package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb+srv://devender:Devender@cluster0.gcq5qxu.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Atlas!")

	coll := client.Database("practiceDB").Collection("items")
	res, err := coll.InsertOne(ctx, map[string]any{"name": "test2", "createdAt": time.Now()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted ID:", res.InsertedID)
}
