package employee

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Employee3 struct {
	Name string `bson:"name"`
	Dept string `bson:"dept"`
}

func NewEmployee3(name string, dept string) Employee3 {
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
