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

func TestGetAll(t *testing.T) {
	mockRepo := new(mockRepository)
	employee := entity.Employee{
		Name:  "Hossain Mahmud",
		Id:    1,
		Title: "Software Engineer",
		Team:  "Stash",
		Email: "hmsayem@gmail.com",
	}
	mockRepo.On("GetAll").Return([]entity.Employee{employee}, nil)
	testService := NewEmployeeService(mockRepo)
	result, err := testService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, result[0].Id)
	assert.Equal(t, "Hossain Mahmud", result[0].Name)
	assert.Equal(t, "Software Engineer", result[0].Title)
	assert.Equal(t, "Stash", result[0].Team)
	assert.Equal(t, "hmsayem@gmail.com", result[0].Email)
	mockRepo.AssertExpectations(t)
}

func TestGetEmployeeByID(t *testing.T) {
	mockRepo := new(mockRepository)
	employee := &entity.Employee{
		Name:  "Hossain Mahmud",
		Id:    1,
		Title: "Software Engineer",
		Team:  "Stash",
		Email: "hmsayem@gmail.com",
	}
	mockRepo.On("GetEmployeeByID").Return(employee, nil)
	testService := NewEmployeeService(mockRepo)
	result, err := testService.Get("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Hossain Mahmud", result.Name)
	assert.Equal(t, "Software Engineer", result.Title)
	assert.Equal(t, "Stash", result.Team)
	assert.Equal(t, "hmsayem@gmail.com", result.Email)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)
	employee := &entity.Employee{
		Name:  "Hossain Mahmud",
		Id:    1,
		Title: "Software Engineer",
		Team:  "Stash",
		Email: "hmsayem@gmail.com",
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
		Title: "Software Engineer",
		Team:  "Stash",
		Email: "hmsayem@gmail.com",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Name`", err.Error())
}

func TestValidateEmptyEmployeeTitle(t *testing.T) {
	employee := &entity.Employee{
		Id:    1,
		Name:  "Hossain Mahmud",
		Team:  "Stash",
		Email: "hmsayem@gmail.com",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Title`", err.Error())
}

func TestValidateEmptyEmployeeTeam(t *testing.T) {
	employee := &entity.Employee{
		Id:    1,
		Name:  "Hossain Mahmud",
		Title: "Software Engineer",
		Email: "hmsayem@gmail.com",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Team`", err.Error())
}

func TestValidateEmptyEmployeeEmail(t *testing.T) {
	employee := &entity.Employee{
		Id:    1,
		Name:  "Hossain Mahmud",
		Title: "Software Engineer",
		Team:  "Stash",
	}
	testService := NewEmployeeService(nil)
	err := testService.Validate(employee)
	assert.NotNil(t, err)
	assert.Equal(t, "empty field `Email`", err.Error())
}
