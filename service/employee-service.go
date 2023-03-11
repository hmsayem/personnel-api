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
	Delete(id string) error
	Validate(employee *entity.Employee) error
}

type employee struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employee{
		repo: repo,
	}
}

func (e *employee) GetAll() ([]entity.Employee, error) {
	return e.repo.GetAll()
}

func (e *employee) Get(id string) (*entity.Employee, error) {
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return e.repo.Get(employeeId)
}

func (e *employee) Update(id string, employee *entity.Employee) error {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return e.repo.Update(parsedId, employee)
}

func (e *employee) Create(employee *entity.Employee) error {
	rand.Seed(time.Now().UnixNano())
	employee.Id = rand.Intn(1000)
	return e.repo.Save(employee)
}

func (e *employee) Delete(id string) error {
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return e.repo.Delete(employeeId)
}

func (*employee) Validate(employee *entity.Employee) error {
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
