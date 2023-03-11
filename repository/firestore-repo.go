package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/hmsayem/clean-architecture-implementation/entity"
	"google.golang.org/api/iterator"
)

const (
	projectId      = "employee-server"
	collectionName = "employees"
)

type firestoreRepo struct {
	client *firestore.Client
}

func NewFirestoreRepository() (EmployeeRepository, error) {
	client, err := firestore.NewClient(context.Background(), projectId)
	if err != nil {
		return nil, fmt.Errorf("creating firestore client failed: %v", err)
	}
	return &firestoreRepo{client: client}, nil
}

func (f *firestoreRepo) GetAll() ([]entity.Employee, error) {
	iter := f.client.Collection(collectionName).Documents(context.Background())
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

func (f *firestoreRepo) Get(id int) (*entity.Employee, error) {
	query := f.client.Collection(collectionName).Where("Id", "==", id).Limit(1)
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

func (f *firestoreRepo) Update(id int, employee *entity.Employee) error {
	query := f.client.Collection(collectionName).Where("Id", "==", id).Limit(1)

	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return fmt.Errorf("querying firestore documents failed: %v", err)
	}
	if len(docs) == 0 {
		return nil
	}
	docRef := docs[0].Ref

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

	_, err = docRef.Update(context.Background(), updates)
	if err != nil {
		return fmt.Errorf("updating firestore document failed: %v", err)
	}
	return nil
}

func (f *firestoreRepo) Save(employee *entity.Employee) error {
	_, _, err := f.client.Collection(collectionName).Add(context.Background(), map[string]interface{}{
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

func (f *firestoreRepo) Delete(id int) error {
	query := f.client.Collection(collectionName).Where("Id", "==", id).Limit(1)

	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return fmt.Errorf("querying firestore documents failed: %v", err)
	}
	if len(docs) == 0 {
		return nil
	}
	docRef := docs[0].Ref

	_, err = docRef.Delete(context.Background())
	if err != nil {
		return fmt.Errorf("deleting firestore document failed: %v", err)
	}
	return nil
}
