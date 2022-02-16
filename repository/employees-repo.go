package repository

import (
	"github.com/hmsayem/clean-architecture-implementation/entity"
)

type EmployeeRepository interface {
	Save(employee *entity.Employee) error
	GetAll() ([]entity.Employee, error)
}
