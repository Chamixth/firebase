package dao

import (
	"context"
	"firebase_generated/dbConfig"
	
)

func DeleteUser(ctx context.Context, userName string) error {
	query := dbConfig.DATABASE_FireBase.Collection("Users").Where("Username", "==", userName).Limit(1)
	docSnap, err := query.Documents(ctx).Next()

	if err != nil {
		return err
	}
	docID := docSnap.Ref.ID
	_, err = dbConfig.DATABASE_FireBase.Collection("User").Doc(docID).Delete(ctx)

	if err != nil {
		return err
	}
	return nil
}