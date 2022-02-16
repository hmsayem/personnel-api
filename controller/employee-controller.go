package controller

import (
	"encoding/json"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"github.com/hmsayem/clean-architecture-implementation/errors"
	"github.com/hmsayem/clean-architecture-implementation/service"
	"log"
	"net/http"
)

type EmployeeController interface {
	GetEmployees(writer http.ResponseWriter, request *http.Request)
	AddEmployee(writer http.ResponseWriter, request *http.Request)
}

type controller struct{}

var (
	employeeService service.EmployeeService
)

func NewEmployeeController(service service.EmployeeService) EmployeeController {
	employeeService = service
	return &controller{}
}

func (*controller) GetEmployees(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	employees, err := employeeService.GetAll()
	if err != nil {
		log.Printf("getting employees failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to get the employees"}); err != nil {
			return
		}
		return
	}
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(employees); err != nil {
		return
	}
}

func (*controller) AddEmployee(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	var employee entity.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		log.Printf("unmarshalling data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to add the new employee"}); err != nil {
			return
		}
		return
	}
	err = employeeService.Validate(&employee)
	if err != nil {
		log.Printf("validating data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: err.Error()}); err != nil {
			return
		}
		return
	}
	err = employeeService.Create(&employee)
	if err != nil {
		log.Printf("saving data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to add the new employee"}); err != nil {
			return
		}
		return
	}
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(employee); err != nil {
		return
	}
}
