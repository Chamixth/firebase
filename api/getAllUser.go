package api

import (
	"context"
	"encoding/json"
	"firebase_generated/dao"

	"github.com/aws/aws-lambda-go/events"
)

func GetAllUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	result, err := dao.DB_FindallUser(ctx)

	if err != nil {

		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to retrieve user"}, err
	}
	responseBody, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}