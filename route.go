package main

import (
	"encoding/json"
	"github.com/hmsayem/employee-server/entity"
	"github.com/hmsayem/employee-server/repository"
	"math/rand"
	"net/http"
)

var (
	repo = repository.NewEmployeeRepository()
)

func init() {

}
func getEmployees(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	employees, err := repo.FindAll()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "failed to get the employees"}`))
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(employees)
}
func addEmployee(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	var employee entity.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "failed to add new employee"}`))
		return
	}
	employee.Id = rand.Int63()
	err = repo.Save(&employee)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "failed to add new employee"}`))
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(employee)

}
