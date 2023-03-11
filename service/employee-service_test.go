package service

import (
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (repo *mockRepository) GetAll() ([]entity.Employee, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.([]entity.Employee), args.Error(1)
}
func (repo *mockRepository) Save(employee *entity.Employee) error {
	args := repo.Called()
	return args.Error(0)
}

func (repo *mockRepository) Get(id int) (*entity.Employee, error) {
	args := repo.Called()
	result := args.Get(0)
	return result.(*entity.Employee), args.Error(1)
}

func (repo *mockRepository) Update(id int, employee *entity.Employee) error {
	args := repo.Called()
	return args.Error(0)
}

func (repo *mockRepository) Delete(id int) error {
	args := repo.Called()
	return args.Error(0)
}
func TestGetAll(t *testing.T) {
	mockRepo := new(mockRepository)
	employees := []entity.Employee{
		{
			Name:  "n1",
			Id:    1,
			Title: "t1",
			Team:  "s1",
			Email: "n1@gmail.com",
		},
		{
			Name:  "n2",
			Id:    2,
			Title: "t2",
			Team:  "s2",
			Email: "n2@gmail.com",
		},
	}
	mockRepo.On("GetAll").Return(employees, nil)
	testService := NewEmployeeService(mockRepo)
	result, err := testService.GetAll()
	assert.Nil(t, err)

	for i := range result {
		assert.Equal(t, employees[i].Id, result[i].Id)
		assert.Equal(t, employees[i].Name, result[i].Name)
		assert.Equal(t, employees[i].Title, result[i].Title)
		assert.Equal(t, employees[i].Team, result[i].Team)
		assert.Equal(t, employees[i].Email, result[i].Email)
	}

	mockRepo.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	mockRepo := new(mockRepository)
	employee := &entity.Employee{
		Name:  "n1",
		Id:    1,
		Title: "t1",
		Team:  "s1",
		Email: "n1@gmail.com",
	}
	mockRepo.On("Get").Return(employee, nil)
	testService := NewEmployeeService(mockRepo)
	result, err := testService.Get("1")
	assert.Nil(t, err)
	assert.Equal(t, employee.Id, result.Id)
	assert.Equal(t, employee.Name, result.Name)
	assert.Equal(t, employee.Title, result.Title)
	assert.Equal(t, employee.Team, result.Team)
	assert.Equal(t, employee.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)
	employee := &entity.Employee{
		Name:  "n1",
		Id:    1,
		Title: "t1",
		Team:  "s1",
		Email: "n1@gmail.com",
	}
	mockRepo.On("Save").Return(nil)
	testService := NewEmployeeService(mockRepo)
	err := testService.Create(employee)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestValidateEmptyEmployee(t *testing.T) {
	testService := NewEmployeeService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "employee is empty", err.Error())
}

func TestValidateEmptyEmployeeName(t *testing.T) {
	employee := &entity.Employee{
		Id:    1,
		Title: "t1",
		Team:  "s1",
		Email: "n1@gmail.com",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Name`", err.Error())
}

func TestValidateEmptyEmployeeTitle(t *testing.T) {
	employee := &entity.Employee{
		Name:  "n1",
		Id:    1,
		Team:  "s1",
		Email: "n1@gmail.com",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Title`", err.Error())
}

func TestValidateEmptyEmployeeTeam(t *testing.T) {
	employee := &entity.Employee{
		Name:  "n1",
		Id:    1,
		Title: "t1",
		Email: "n1@gmail.com",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Team`", err.Error())
}

func TestValidateEmptyEmployeeEmail(t *testing.T) {
	employee := &entity.Employee{
		Name:  "n1",
		Id:    1,
		Title: "t1",
		Team:  "s1",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Email`", err.Error())
}
