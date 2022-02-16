package service

import (
	"errors"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"github.com/hmsayem/clean-architecture-implementation/repository"
	"math/rand"
)

type EmployeeService interface {
	Validate(employee *entity.Employee) error
	Create(employee *entity.Employee) error
	GetAll() ([]entity.Employee, error)
}

type service struct{}

var (
	employeeRepo repository.EmployeeRepository
)

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	employeeRepo = repo
	return &service{}
}

func (*service) Validate(employee *entity.Employee) error {
	if employee == nil {
		return errors.New("employee is empty")
	}
	if employee.Name == "" {
		return errors.New("empty field `Name`")
	}
	if employee.Title == "" {
		return errors.New("empty field `Title`")
	}
	if employee.Team == "" {
		return errors.New("empty field `Team`")
	}
	if employee.Email == "" {
		return errors.New("empty field `Email`")
	}
	return nil
}

func (*service) Create(employee *entity.Employee) error {
	employee.Id = rand.Int63()
	return employeeRepo.Save(employee)
}

func (*service) GetAll() ([]entity.Employee, error) {
	return employeeRepo.GetAll()
}
