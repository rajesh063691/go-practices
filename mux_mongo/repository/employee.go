package repository

import (
	"context"
	"fmt"
	"mux_mongo/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// it holds the mongodb	Collection instance
type EmployeeRepo struct {
	MongoDBCollection *mongo.Collection
}

// this will return inserted mongodb id or error if any
func (r *EmployeeRepo) CreateEmployee(ctx context.Context, e *models.Employee) (interface{}, error) {
	// create employee
	//db := r.MongoDBCollection.Database().Collection("employee")
	result, err := r.MongoDBCollection.InsertOne(context.Background(), e)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil

}

func (r *EmployeeRepo) GetEmployee(ctx context.Context, empID string) (emp *models.Employee, err error) {
	// get employee
	emp = &models.Employee{}
	filter := bson.M{"employee_id": empID}
	err = r.MongoDBCollection.FindOne(ctx, filter).Decode(&emp)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func (r *EmployeeRepo) FindAllEmployee(ctx context.Context) (emps []models.Employee, err error) {
	// get all employee
	emps = make([]models.Employee, 0)
	filer := bson.M{}
	cursor, err := r.MongoDBCollection.Find(context.Background(), filer)
	if err != nil {
		return emps, fmt.Errorf("error in fetching employee: %v", err)

	}
	defer func() { cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		var emp models.Employee
		err = cursor.Decode(&emp)
		if err != nil {
			return emps, fmt.Errorf("Error in decoding employee: %v", err)
		}
		emps = append(emps, emp)
	}

	return emps, nil

}

func (r *EmployeeRepo) UpdateEmployee(ctx context.Context, empID string, e *models.Employee) (count int64, err error) {
	// update employee
	filter := bson.M{"employee_id": empID}
	update := bson.M{"$set": e}
	result, err := r.MongoDBCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return count, fmt.Errorf("error in updating employee: %v", err)
	}
	count = result.ModifiedCount
	return count, nil

}

func (r *EmployeeRepo) DeleteEmployee(ctx context.Context, empID string) (count int64, err error) {
	// delete employee
	filter := bson.M{"employee_id": empID}
	result, err := r.MongoDBCollection.DeleteOne(ctx, filter)
	if err != nil {
		return count, fmt.Errorf("error in deleting employee: %v", err)
	}
	count = result.DeletedCount
	return count, nil

}

func (r *EmployeeRepo) DeleteAllEmployee(ctx context.Context) (int64, error) {
	result, err := r.MongoDBCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return 0, fmt.Errorf("error in deleting employee: %v", err)
	}

	return result.DeletedCount, nil
}
