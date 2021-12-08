package service

import (
	"errors"
	"github.com/hmsayem/employee-server/entity"
	"github.com/hmsayem/employee-server/repository"
	"math/rand"
)

type EmployeeService interface {
	Validate(employee *entity.Employee) error
	Create(employee *entity.Employee) error
	FindAll() ([]entity.Employee, error)
}

type service struct{}

var (
	employeeRepo repository.EmployeeRepository
	//repo = repository.NewFirestoreRepository()
)

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	employeeRepo = repo
	return &service{}
}

func (*service) Validate(employee *entity.Employee) error {
	if employee == nil {
		err := errors.New("employee is empty")
		return err
	}
	return nil
}

func (*service) Create(employee *entity.Employee) error {
	employee.Id = rand.Int63()
	return employeeRepo.Save(employee)
}

func (*service) FindAll() ([]entity.Employee, error) {
	return employeeRepo.FindAll()
}
