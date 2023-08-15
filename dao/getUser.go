package dao

import (
	"context"
	"firebase_generated/dbConfig"
	"firebase_generated/dto"
)

func GetUser(ctx context.Context, userName string) (*dto.User, error) {
	var request dto.User
	query := dbConfig.DATABASE_FireBase.Collection("Users").Where("Username", "==", userName).Limit(1)
	docSnap, err := query.Documents(ctx).Next()
	if err != nil {
		return nil, err
	}
	err = docSnap.DataTo(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}
