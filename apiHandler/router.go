package apiHandler

import (
	"context"
	"firebase_generated/api"
	"firebase_generated/dbConfig"
	

	"github.com/aws/aws-lambda-go/events"
)

func Router(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {


	////////////////////////////////////////////////////MongoDb//////////////////////////////////////////////////
	dbConfig.ConnectToMongoDB()
	method := request.QueryStringParameters["method"]
	////////////////////////////////////////////////////MongoDb//////////////////////////////////////////////////








	////////////////////////////////////////////////////Firebase-firestore//////////////////////////////////////

	// Connect to Firestore with S3 credentials
	credentialsURL := "https://firebasejson.s3.amazonaws.com/go-lambda-firebase-adminsdk-hpadx-90be10eb70.json"
	dbConfig.ConnectToFirestore(ctx, credentialsURL)
	
	////////////////////////////////////////////////////Firebase-firestore//////////////////////////////////////






	////////////////////////////////////////////////////Common/////////////////////////////////////////////////

	switch method {
	case "createUser":
		return api.CreateUser(ctx, request)
	case "getUser":
		return api.GetUser(ctx, request)
	case "deleteUser":
		return api.DeleteUser(ctx, request)
	case "updateUser":
		return api.UpdateUser(ctx, request)
	case "getAllUser":
		return api.GetAllUser(ctx,request)
	default:
		return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Method Not Allowed"}, nil
	}

	////////////////////////////////////////////////////Common/////////////////////////////////////////////////
}
