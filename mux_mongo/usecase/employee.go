package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"mux_mongo/models"
	"mux_mongo/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoDBCollection *mongo.Collection
}

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (s *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// read the request body
	w.Header().Add("Content-Type", "application/json")
	var emp models.Employee
	res := &APIResponse{}
	defer json.NewEncoder(w).Encode(res)
	// decode the request body
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = uuid.New().String()
	// instantiate the repository
	repo := repository.EmployeeRepo{
		MongoDBCollection: s.MongoDBCollection,
	}

	// create employee
	insertID, err := repo.CreateEmployee(r.Context(), &emp)
	if err != nil {
		log.Printf("Error in creating employee: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}
	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)
	log.Println("Employee created successfully with id: ", insertID, emp)
}
func (s *EmployeeService) GetEmpByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &APIResponse{}
	defer json.NewEncoder(w).Encode(res)
	empID := mux.Vars(r)["id"]
	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = fmt.Sprintf("id provided is empty :: %s", empID)
		return
	}

	repo := repository.EmployeeRepo{
		MongoDBCollection: s.MongoDBCollection,
	}

	emp, err := repo.GetEmployee(r.Context(), empID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return

	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
	log.Println("Employee fetched successfully with id: ", empID)
}

func (s *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{
		MongoDBCollection: s.MongoDBCollection,
	}

	emps, err := repo.FindAllEmployee(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}
	res.Data = emps
	w.WriteHeader(http.StatusOK)
	log.Println("Employee fetched successfully")
}
func (s *EmployeeService) UpdateEmpByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = fmt.Sprintf("id provided is empty :: %s", empID)
		return
	}

	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = empID

	repo := repository.EmployeeRepo{
		MongoDBCollection: s.MongoDBCollection,
	}

	count, err := repo.UpdateEmployee(r.Context(), empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
	log.Println("Employee updated successfully with id: ", empID)
}

func (s *EmployeeService) DeleteEmpByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = fmt.Sprintf("id provided is empty :: %s", empID)
		return
	}

	repo := repository.EmployeeRepo{
		MongoDBCollection: s.MongoDBCollection,
	}

	count, err := repo.DeleteEmployee(r.Context(), empID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
	log.Println("Employee deleted successfully with id: ", empID)
}
func (s *EmployeeService) DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &APIResponse{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{
		MongoDBCollection: s.MongoDBCollection,
	}

	count, err := repo.DeleteAllEmployee(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
	log.Println("All Employee deleted successfully")
}
