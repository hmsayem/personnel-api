package controller

import (
	"encoding/json"
	"github.com/hmsayem/clean-architecture-implementation/cache"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"github.com/hmsayem/clean-architecture-implementation/errors"
	"github.com/hmsayem/clean-architecture-implementation/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type EmployeeController interface {
	GetAll(writer http.ResponseWriter, request *http.Request)
	Get(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Add(writer http.ResponseWriter, request *http.Request)
}

type employee struct {
	service service.EmployeeService
	cache   cache.EmployeeCache
}

func NewEmployeeController(service service.EmployeeService, cache cache.EmployeeCache) EmployeeController {
	return &employee{
		service: service,
		cache:   cache,
	}
}

func (e *employee) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	employees, err := e.service.GetAll()
	if err != nil {
		log.Printf("failed to get employees: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to get employees"}); err != nil {
			return
		}
		return
	}
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(employees); err != nil {
		return
	}
}

func (e *employee) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	id := strings.Split(request.URL.Path, "/")[2]
	employee, err := e.cache.Get(id)
	if err != nil {
		log.Printf("failed to get value from cache: %v", err)
		employee, err = e.service.Get(id)
		if err != nil {
			log.Printf("failed to get employee: %v", err)
			writer.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to get employee"}); err != nil {
				return
			}
			return
		}
		e.updateCache(employee)
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(employee); err != nil {
		return
	}
}

func (e *employee) Add(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	var employee entity.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		log.Printf("unmarshalling data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to add new employee"}); err != nil {
			return
		}
		return
	}
	err = e.service.Validate(&employee)
	if err != nil {
		log.Printf("validating data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: err.Error()}); err != nil {
			return
		}
		return
	}
	err = e.service.Create(&employee)
	if err != nil {
		log.Printf("saving data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to add new employee"}); err != nil {
			return
		}
		return
	}
	e.updateCache(&employee)
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(employee); err != nil {
		return
	}
}

func (e *employee) Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "application/json")
	var employee entity.Employee
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		log.Printf("unmarshalling data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to update employee"}); err != nil {
			return
		}
		return
	}

	id := strings.Split(request.URL.Path, "/")[2]
	err = e.service.Update(id, &employee)
	if err != nil {
		log.Printf("updating data failed: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(writer).Encode(errors.ServiceError{Message: "failed to update employee"}); err != nil {
			return
		}
		return
	}
	if err := e.cache.Del(id); err != nil {
		log.Printf("failed to delete key in cache: %v", err)
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (e *employee) updateCache(employee *entity.Employee) {
	if err := e.cache.Set(strconv.Itoa(employee.Id), employee); err != nil {
		log.Printf("failed to save key in cache: %v", err)
	}
}
