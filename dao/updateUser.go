package dao

import (
	"context"
	"firebase_generated/dbConfig"
	"firebase_generated/dto"
)

func UpdateUser(ctx context.Context, application dto.User) error {
	query := dbConfig.DATABASE_FireBase.Collection("Users").Where("Username", "==", application.Username).Limit(1)
	docSnap, err := query.Documents(ctx).Next()
	if err != nil {
		return err
	}
	docID := docSnap.Ref.ID
	_, err = dbConfig.DATABASE_FireBase.Collection("Users").Doc(docID).Set(ctx, application)
	if err != nil {
		return err
	}
	return err
}