package controller

import (
	"encoding/json"
	"github.com/hmsayem/clean-architecture-implementation/cache"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"github.com/hmsayem/clean-architecture-implementation/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	contentType     = "Content-type"
	applicationJSON = "application/json"

	ErrFailedToGetEmployees   = "failed to get employees"
	ErrFailedToGetEmployee    = "failed to get employee"
	ErrFailedToAddEmployee    = "failed to add new employee"
	ErrFailedToUpdateEmployee = "failed to update employee"
	ErrFailedToDeleteEmployee = "failed to delete employee"
	ErrInvalidEmployeeData    = "invalid employee data"
)

type EmployeeController interface {
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Add(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
}

type controller struct {
	service service.EmployeeService
	cache   cache.EmployeeCache
}

func NewEmployeeController(service service.EmployeeService, cache cache.EmployeeCache) EmployeeController {
	return &controller{
		service: service,
		cache:   cache,
	}
}

func (e *controller) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(contentType, applicationJSON)
	employees, err := e.service.GetAll()
	if err != nil {
		log.Printf("failed to get employees: %v", err)
		http.Error(writer, ErrFailedToGetEmployees, http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(employees); err != nil {
		log.Printf("failed to encode employees: %v", err)
		return
	}
}

func (e *controller) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(contentType, applicationJSON)
	id := strings.Split(request.URL.Path, "/")[2]
	employee, err := e.cache.Get(id)
	if err != nil {
		log.Printf("failed to get value from cache: %v", err)
		employee, err = e.service.Get(id)
		if err != nil {
			log.Printf("failed to get employee: %v", err)
			http.Error(writer, ErrFailedToGetEmployee, http.StatusInternalServerError)
			return
		}
		e.updateCache(employee)
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(employee); err != nil {
		log.Printf("failed to encode employee: %v", err)
		return
	}
}

func (e *controller) Add(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(contentType, applicationJSON)
	var employee entity.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		log.Printf("failed to unmarshal data: %v", err)
		http.Error(writer, ErrFailedToAddEmployee, http.StatusInternalServerError)
		return
	}
	err = e.service.Validate(&employee)
	if err != nil {
		log.Printf("invalid employee data: %v", err)
		http.Error(writer, ErrInvalidEmployeeData, http.StatusInternalServerError)
		return
	}
	err = e.service.Create(&employee)
	if err != nil {
		log.Printf("failed to add employee: %v", err)
		http.Error(writer, ErrFailedToAddEmployee, http.StatusInternalServerError)
		return
	}
	e.updateCache(&employee)
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(employee); err != nil {
		log.Printf("failed to encode employee: %v", err)
		return
	}
}

func (e *controller) Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(contentType, applicationJSON)
	var employee entity.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		log.Printf("failed to unmarshal data: %v", err)
		http.Error(writer, ErrFailedToUpdateEmployee, http.StatusInternalServerError)
		return
	}

	id := strings.Split(request.URL.Path, "/")[2]
	err = e.service.Update(id, &employee)
	if err != nil {
		log.Printf("failed to update employee: %v", err)
		http.Error(writer, ErrFailedToUpdateEmployee, http.StatusInternalServerError)
		return
	}
	if err := e.cache.Del(id); err != nil {
		log.Printf("failed to delete key in cache: %v", err)
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (e *controller) updateCache(employee *entity.Employee) {
	if err := e.cache.Set(strconv.Itoa(employee.Id), employee); err != nil {
		log.Printf("failed to save key in cache: %v", err)
	}
}

func (e *controller) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(contentType, applicationJSON)
	id := strings.Split(request.URL.Path, "/")[2]

	if err := e.service.Delete(id); err != nil {
		log.Printf("failed to delete employee: %v", err)
		http.Error(writer, ErrFailedToDeleteEmployee, http.StatusInternalServerError)
		return
	}
	if err := e.cache.Del(id); err != nil {
		log.Printf("failed to delete key in cache: %v", err)
	}
	writer.WriteHeader(http.StatusNoContent)
}
