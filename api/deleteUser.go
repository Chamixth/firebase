package api
import (
	"context"
	"encoding/json"
	"fmt"
	"firebase_generated/dao"
	

	"github.com/aws/aws-lambda-go/events"
)

func DeleteUser(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userName := event.QueryStringParameters["userName"]

	err := dao.DeleteUser(ctx, userName)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to delete user"}, err
	}

	responseBody := fmt.Sprintf("User '%s' deleted successfully", userName)
	responseBodyJSON, _ := json.Marshal(responseBody)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBodyJSON),
	}, nil
}
