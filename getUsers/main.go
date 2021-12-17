package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Array of user where we are going to store the fetched users
var arr []User

// Lambda function that fetches all users from DynamoDB
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse,
	error) {
	// Fetch user is the controller dedicated to fetch the user
	arr, err := FetchUsers()
	if err != nil {
		json := fmt.Sprintf(`{
			"message": "there was an error at the server"
		}`)
		return response(json, http.StatusInternalServerError), nil
	}

	// If there are no errors but it could not find any user.
	if len(arr) == 0 {
		json := fmt.Sprintf(`{
			"message": "the database is empty",
		}`)
		return response(json, 200), nil
	}

	arrJson, _ := json.Marshal(arr)

	// If we find a user, send it through the response body.
	json := fmt.Sprintf(`{
			"message": "users fetched succesfully",
			"user" : %v
		}`, string(arrJson))
	return events.APIGatewayProxyResponse{Body: json, StatusCode: 200}, nil

}

func main() {
	lambda.Start(handleRequest)
}

// Send HTTP response
func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}

// GOOS=linux GOARCH=amd64 go build -o main main.go connection.go model.go
// zip main.zip main
