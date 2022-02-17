package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/hmsayem/clean-architecture-implementation/entity"
)

const (
	projectId      = "employee-server"
	collectionName = "employees"
)

type firestoreRepo struct{}

func NewFirestoreRepository() EmployeeRepository {
	return &firestoreRepo{}
}

func (*firestoreRepo) Save(employee *entity.Employee) error {

	client, err := getClient()
	if err != nil {
		return err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(context.Background(), map[string]interface{}{
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

func (*firestoreRepo) GetAll() ([]entity.Employee, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	docs, err := getDocs(client)
	if err != nil {
		return nil, err
	}
	var employees []entity.Employee
	for _, doc := range docs {
		employee := entity.Employee{
			Id:    int(doc.Data()["Id"].(int64)),
			Name:  doc.Data()["Name"].(string),
			Title: doc.Data()["Title"].(string),
			Team:  doc.Data()["Team"].(string),
			Email: doc.Data()["Email"].(string),
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (*firestoreRepo) GetEmployee(id int) (*entity.Employee, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	docs, err := getDocs(client)
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		if int(doc.Data()["Id"].(int64)) == id {
			return &entity.Employee{
				Id:    int(doc.Data()["Id"].(int64)),
				Name:  doc.Data()["Name"].(string),
				Title: doc.Data()["Title"].(string),
				Team:  doc.Data()["Team"].(string),
				Email: doc.Data()["Email"].(string),
			}, nil
		}
	}
	return nil, nil
}

func getClient() (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("creating firestore client failed: %v", err)
	}
	return client, nil
}

func getDocs(client *firestore.Client) ([]*firestore.DocumentSnapshot, error) {
	iterator := client.Collection(collectionName).Documents(context.Background())
	docs, err := iterator.GetAll()
	if err != nil {
		return nil, fmt.Errorf("iterating firestore documents failed: %v", err)
	}
	return docs, nil
}
