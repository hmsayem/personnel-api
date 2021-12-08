package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/hmsayem/employee-server/entity"
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
		return fmt.Errorf("creating firestore client failed: %v", err)
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
		return fmt.Errorf("creating firestore document failed: %v", err)
	}
	return nil
}

func (*firestoreRepo) FindAll() ([]entity.Employee, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("creating firestore client failed: %v", err)
	}
	defer client.Close()
	var employees []entity.Employee
	iterator := client.Collection(collectionName).Documents(ctx)
	docs, err := iterator.GetAll()
	if err != nil {
		return nil, fmt.Errorf("iterating firestore documents failed: %v", err)
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
