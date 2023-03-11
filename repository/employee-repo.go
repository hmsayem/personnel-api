package repository

import (
	"github.com/hmsayem/clean-architecture-implementation/entity"
)

type EmployeeRepository interface {
	GetAll() ([]entity.Employee, error)
	Get(id int) (*entity.Employee, error)
	Update(id int, employee *entity.Employee) error
	Save(employee *entity.Employee) error
	Delete(id int) error
}
