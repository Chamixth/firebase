package main

import (
	"firebase_generated/apiHandler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	//lambda.Start(apiHandler.Router)asdasdasdasdasdasd
	lambda.Start(apiHandler.Router)
}
