package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/hmsayem/employee-server/entity"
	"log"
)

const (
	projectId      = "employee-server"
	collectionName = "employees"
)

type firestoreRepo struct {
}

func NewFirestoreRepository() EmployeeRepository {
	return &firestoreRepo{}
}

func (*firestoreRepo) Save(employee *entity.Employee) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create Firestore client: %v", err)
		return err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":    employee.Id,
		"Name":  employee.Name,
		"Title": employee.Title,
		"Team":  employee.Team,
		"Email": employee.Email,
	})
	if err != nil {
		log.Printf("Failed to add new employee: %v", err)
		return err
	}
	return nil
}

func (*firestoreRepo) FindAll() ([]entity.Employee, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var employees []entity.Employee
	iterator := client.Collection(collectionName).Documents(ctx)
	docs, err := iterator.GetAll()
	if err != nil {
		log.Printf("Failed to get the employee list: %v", err)
		return nil, err
	}
	for _, doc := range docs {
		employee := entity.Employee{
			Id:    doc.Data()["Id"].(int64),
			Name:  doc.Data()["Name"].(string),
			Title: doc.Data()["Title"].(string),
			Team:  doc.Data()["Team"].(string),
			Email: doc.Data()["Email"].(string),
		}
		employees = append(employees, employee)
	}
	return employees, nil
}
