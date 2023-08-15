package api

import (
	"context"
	"encoding/json"
	"firebase_generated/dao"
	

	"github.com/aws/aws-lambda-go/events"
)

func GetUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	username := event.QueryStringParameters["userName"]

	result, err := dao.GetUser(ctx, username)

	if err != nil {

		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to retrieve user"}, err
	}
	responseBody, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}