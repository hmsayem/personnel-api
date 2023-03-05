package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
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

	iter := client.Collection(collectionName).Documents(context.Background())
	docs, err := iter.GetAll()
	if err != nil {
		return nil, fmt.Errorf("iterating firestore documents failed: %v", err)
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

func (repo *firestoreRepo) Get(id int) (*entity.Employee, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	query := client.Collection(collectionName).Where("Id", "==", id).Limit(1)
	iter := query.Documents(context.Background())
	docSnapshot, err := iter.Next()
	if err != nil {
		if err == iterator.Done {
			return nil, nil // Employee not found
		}
		return nil, fmt.Errorf("iterating firestore documents failed: %v", err)
	}

	return &entity.Employee{
		Id:    int(docSnapshot.Data()["Id"].(int64)),
		Name:  docSnapshot.Data()["Name"].(string),
		Title: docSnapshot.Data()["Title"].(string),
		Team:  docSnapshot.Data()["Team"].(string),
		Email: docSnapshot.Data()["Email"].(string),
	}, nil
}

func (repo *firestoreRepo) Update(id int, employee *entity.Employee) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// Get the document reference for the employee with the given ID.
	docRef := client.Collection(collectionName).Doc(strconv.Itoa(id))

	// Check if the document exists.
	_, err = docRef.Get(context.Background())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil
		}
		return fmt.Errorf("getting firestore document failed: %v", err)
	}

	// Construct a slice of Firestore update structs.
	var updates []firestore.Update
	if employee.Name != "" {
		updates = append(updates, firestore.Update{Path: "Name", Value: employee.Name})
	}
	if employee.Title != "" {
		updates = append(updates, firestore.Update{Path: "Title", Value: employee.Title})
	}
	if employee.Team != "" {
		updates = append(updates, firestore.Update{Path: "Team", Value: employee.Team})
	}
	if employee.Email != "" {
		updates = append(updates, firestore.Update{Path: "Email", Value: employee.Email})
	}

	// Update the fields of the document.
	_, err = docRef.Update(context.Background(), updates)
	if err != nil {
		return fmt.Errorf("updating firestore document failed: %v", err)
	}

	return nil
}

func getClient() (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("creating firestore client failed: %v", err)
	}
	return client, nil
}
