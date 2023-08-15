package api

import (
	"context"
	"encoding/json"
	"fmt"
	"firebase_generated/dao"
	"firebase_generated/dto"
	

	"github.com/aws/aws-lambda-go/events"
)

func UpdateUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var updatedUser dto.User
	err := json.Unmarshal([]byte(event.Body), &updatedUser)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	}
	if err := dao.UpdateUser(ctx, updatedUser); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to update user"}, err
	}
	responseBody := fmt.Sprintf("User '%s' updated successfully", updatedUser.Username)
	responseBodyJSON, _ := json.Marshal(responseBody)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBodyJSON),
	}, nil
}