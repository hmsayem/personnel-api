package repository

import (
	"github.com/hmsayem/clean-architecture-implementation/entity"
)

type EmployeeRepository interface {
	GetAll() ([]entity.Employee, error)
	GetEmployeeByID(id int) (*entity.Employee, error)
	Save(employee *entity.Employee) error
}
