package api

import (
	"context"
	"encoding/json"
	"fmt"
	"firebase_generated/dao"
	"firebase_generated/dto"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func CreateUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user dto.User

	err := json.Unmarshal([]byte(event.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "Invalid request body"}, nil
	}
	if err := dao.CreateUser(ctx, user); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Failed to add user"}, err
	}
	

	responseBody := fmt.Sprintf("User '%s' added successfully", user.Username)
	responseBodyJSON, _ := json.Marshal(responseBody)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBodyJSON),
	}, nil
}