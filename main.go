package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/devender81/mongodb/employee"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	emp3 := employee.NewEmployee3("adv1", "devops")
	fmt.Println(emp3)

	uri := "mongodb+srv://devender:Devender@cluster0.gcq5qxu.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	coll := client.Database("practiceDB").Collection("employees")

	id, err := employee.InsertEmployee(ctx, coll, emp3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted employee with ID:", id)
}
