package repository

import (
	"github.com/hmsayem/employee-server/entity"
)

type EmployeeRepository interface {
	Save(employee *entity.Employee) error
	FindAll() ([]entity.Employee, error)
}
