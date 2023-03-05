package service

import (
	"errors"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"github.com/hmsayem/clean-architecture-implementation/repository"
	"math/rand"
	"strconv"
	"time"
)

type EmployeeService interface {
	GetAll() ([]entity.Employee, error)
	Get(id string) (*entity.Employee, error)
	Update(id string, employee *entity.Employee) error
	Create(employee *entity.Employee) error
	Validate(employee *entity.Employee) error
}

type service struct{}

var (
	employeeRepo repository.EmployeeRepository
)

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	employeeRepo = repo
	return &service{}
}

func (*service) GetAll() ([]entity.Employee, error) {
	return employeeRepo.GetAll()
}

func (*service) Get(id string) (*entity.Employee, error) {
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return employeeRepo.Get(employeeId)
}

func (*service) Update(id string, employee *entity.Employee) error {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return employeeRepo.Update(parsedId, employee)
}

func (*service) Create(employee *entity.Employee) error {
	rand.Seed(time.Now().UnixNano())
	employee.Id = rand.Intn(1000)
	return employeeRepo.Save(employee)
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
