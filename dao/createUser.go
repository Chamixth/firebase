package dao

import (
	"context"
	"firebase_generated/dbConfig"
	"firebase_generated/dto"
)


func CreateUser(ctx context.Context, application dto.User) error {
	_, _, err := dbConfig.DATABASE_FireBase.Collection("Users").Add(ctx, application)
	if err != nil {
		return err
	}
	return nil
}