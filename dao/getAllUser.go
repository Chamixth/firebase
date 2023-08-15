package dao

import (
	"firebase_generated/dbConfig"
	"firebase_generated/dto"
    "context"
    "errors"
	"google.golang.org/api/iterator"
)

func DB_FindallUser (ctx context.Context) (*[]dto.User, error) {
	var objects []dto.User

	// Get the collection reference
	docRef := dbConfig.DATABASE_FireBase.Collection("Users")

	// Retrieve all documents
	iter := docRef.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var model dto.User
		if err := doc.DataTo(&model); err != nil {
			errors.New("Error converting document to User struct")
			continue
		}

		objects = append(objects, model)
	}
	return &objects, nil
}

