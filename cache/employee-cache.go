package cache

import "github.com/hmsayem/clean-architecture-implementation/entity"

type EmployeeCache interface {
	Set(key string, value *entity.Employee) error
	Get(key string) (*entity.Employee, error)
	Delete(key string) error
}
