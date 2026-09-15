package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/devender81/mongodb/employee"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	uri := "mongodb+srv://devender:Devender@cluster0.gcq5qxu.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	//uri := "mongodb://localhost:27017"    //connecting mongodb through docker

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

	// ---------- CREATE ----------
	emp := employee.NewEmployee("adv1", "devops")
	insertedID, err := employee.InsertEmployee(ctx, coll, emp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted employee with ID:", insertedID)

	// Cast the returned ID (interface{}) into a bson.ObjectID so we can use it below.
	objID, ok := insertedID.(bson.ObjectID)
	if !ok {
		log.Fatal("could not cast inserted ID to bson.ObjectID")
	}

	// Insert one more so Read-all/Read-by-dept have more than one document to show.
	emp2 := employee.NewEmployee("adv2", "sales")
	if _, err := employee.InsertEmployee(ctx, coll, emp2); err != nil {
		log.Fatal(err)
	}

	// ---------- READ ----------
	fetched, err := employee.GetEmployeeByID(ctx, coll, objID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched by ID:", *fetched)

	all, err := employee.GetAllEmployees(ctx, coll)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All employees:", all)

	devopsFolks, err := employee.GetEmployeesByDept(ctx, coll, "devops")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Devops employees:", devopsFolks)

	// ---------- UPDATE ----------
	modifiedCount, err := employee.UpdateEmployeeDept(ctx, coll, objID, "platform-engineering")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documents modified:", modifiedCount)

	updated, err := employee.GetEmployeeByID(ctx, coll, objID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After update:", *updated)

	// ---------- DELETE ----------
	deletedCount, err := employee.DeleteEmployee(ctx, coll, objID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documents deleted:", deletedCount)

	remaining, err := employee.GetAllEmployees(ctx, coll)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Remaining employees:", remaining)

}
