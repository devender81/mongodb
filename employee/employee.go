package employee

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Employee3 struct {
	ID   bson.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name"`
	Dept string        `bson:"dept"`
}

func NewEmployee(name string, dept string) Employee3 {
	return Employee3{Name: name, Dept: dept}
}

// InsertEmployee inserts an Employee3 into the given collection and returns the inserted ID.
func InsertEmployee(ctx context.Context, coll *mongo.Collection, emp Employee3) (interface{}, error) {
	res, err := coll.InsertOne(ctx, emp)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

// GetEmployeeByID finds a single employee by its ObjectID.
func GetEmployeeByID(ctx context.Context, coll *mongo.Collection, id bson.ObjectID) (*Employee3, error) {
	var emp Employee3
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&emp)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

// GetAllEmployees returns every employee document in the collection.
func GetAllEmployees(ctx context.Context, coll *mongo.Collection) ([]Employee3, error) {
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var employees []Employee3
	if err := cursor.All(ctx, &employees); err != nil {
		return nil, err
	}
	return employees, nil
}

// GetEmployeesByDept finds all employees in a given department.
func GetEmployeesByDept(ctx context.Context, coll *mongo.Collection, dept string) ([]Employee3, error) {
	cursor, err := coll.Find(ctx, bson.M{"dept": dept})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var employees []Employee3
	if err := cursor.All(ctx, &employees); err != nil {
		return nil, err
	}
	return employees, nil
}

// UpdateEmployeeDept updates the department field for a given employee ID.
// Returns the number of documents modified.
func UpdateEmployeeDept(ctx context.Context, coll *mongo.Collection, id bson.ObjectID, newDept string) (int64, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"dept": newDept}}

	res, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

// DeleteEmployee removes a single employee by ID.
// Returns the number of documents deleted (0 or 1).
func DeleteEmployee(ctx context.Context, coll *mongo.Collection, id bson.ObjectID) (int64, error) {
	res, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// DeleteEmployeesByDept removes all employees in a given department.
// Returns the number of documents deleted.
func DeleteEmployeesByDept(ctx context.Context, coll *mongo.Collection, dept string) (int64, error) {
	res, err := coll.DeleteMany(ctx, bson.M{"dept": dept})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
