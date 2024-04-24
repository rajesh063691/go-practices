package repository

import (
	"context"
	"log"
	"mux_mongo/models"
	"strings"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMOngoClient() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://rkyengg009:Scbaqqd0Sxkq8rkM@cluster0.klov5vd.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		log.Fatal("error in connecting to mongodb: ", err)
	}

	log.Println("Connected to MongoDB!")

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed: ", err)
	}

	log.Println("Pinged successfully!")

	return client
}

func TestMongoOperations(t *testing.T) {
	client := NewMOngoClient()
	defer client.Disconnect(context.Background())

	emp1 := uuid.New().String()
	emp1 = strings.Replace(emp1, "-", "", -1)
	emp2 := uuid.New().String()
	emp2 = strings.Replace(emp2, "-", "", -1)

	// create collection
	coll := client.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{
		MongoDBCollection: coll,
	}

	// insert employee data
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := models.Employee{
			EmployeeID: emp1,
			Name:       "Rajesh Kumar",
			Department: "Software Engineer",
		}

		id, err := empRepo.CreateEmployee(context.Background(), &emp)
		if err != nil {
			t.Errorf("Error in inserting employee: %v", err)
		}
		t.Log("insert successful: ", id)
	})

	t.Run("Insert Employee 2", func(t *testing.T) {
		emp := models.Employee{
			EmployeeID: emp2,
			Name:       "Puja Kumair",
			Department: "Banking",
		}

		id, err := empRepo.CreateEmployee(context.Background(), &emp)
		if err != nil {
			t.Errorf("Error in inserting employee: %v", err)
		}
		t.Log("insert successful: ", id)
	})

	// get employee data
	t.Run("Get Employee 1", func(t *testing.T) {
		emp, err := empRepo.GetEmployee(context.Background(), emp1)
		if err != nil {
			t.Errorf("Error in getting employee: %v", err)
		}
		t.Log("Employee data: ", emp)
	})

	// get All employee data
	t.Run("Get All Employees", func(t *testing.T) {
		emp, err := empRepo.FindAllEmployee(context.Background())
		if err != nil {
			t.Errorf("Error in getting all employees: %v", err)
		}
		t.Log("All Employee data: ", emp)
	})

	// update employee data
	t.Run("Update Employee 1", func(t *testing.T) {
		newEmp := models.Employee{
			EmployeeID: emp1,
			Name:       "Rajesh Kumar Yadav",
			Department: "Software Engineering",
		}

		count, err := empRepo.UpdateEmployee(context.Background(), emp1, &newEmp)
		if err != nil {
			t.Errorf("Error in updating employee: %v", err)
		}
		t.Log("Employee updated: ", count)
	})

	// check updated employee data
	t.Run("Get Employee 1 after update", func(t *testing.T) {
		emp, err := empRepo.GetEmployee(context.Background(), emp1)
		if err != nil {
			t.Errorf("Error in getting employee: %v", err)
		}
		t.Log("Updated Employee Name: ", emp.Name)
	})

	// delete employee data
	t.Run("Delete Employee 1", func(t *testing.T) {
		count, err := empRepo.DeleteEmployee(context.Background(), emp1)
		if err != nil {
			t.Errorf("Error in deleting employee: %v", err)
		}
		t.Log("Employee deleted: ", count)
	})

	// get All employee data after delete
	t.Run("Get All Employees after delete", func(t *testing.T) {
		emp, err := empRepo.FindAllEmployee(context.Background())
		if err != nil {
			t.Errorf("Error in getting all employees: %v", err)
		}
		t.Log("All Employee data: ", emp)
	})

	// delete All employee data
	t.Run("Delete All Employees", func(t *testing.T) {
		_, err := empRepo.DeleteAllEmployee(context.Background())
		if err != nil {
			t.Errorf("Error in deleting all employees: %v", err)
		}
		t.Log("All Employees deleted")
	})
}
