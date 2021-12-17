package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// User object where we are going to store the fetched user values
var u User

// Lambda function that fetches a user from DynamoDB with provided ID
//
// Client must send the "id" through the urlparams
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse,
	error) {

	// Get id from url parameters
	id := request.QueryStringParameters["id"]

	// Fetch user is the controller dedicated to fetch the user
	u, err := FetchUser(id)
	if err != nil {
		json := fmt.Sprintf(`{
			"message": "there was an error at the server",
			"provided_id": "%v"
		}`, id)
		return response(json, http.StatusInternalServerError), nil
	}

	// If there are no errors but it could not find any user.
	if u.Username == "" {
		json := fmt.Sprintf(`{
			"message": "could not fetch any user with provided id",
			"provided_id": "%v"
		}`, id)
		return response(json, 200), nil
	}

	// If we find a user, send it through the response body.
	json := fmt.Sprintf(`{
			"message": "user fetched succesfully",
			"user" : {
				"id": "%v",
				"username": "%v",
				"email": "%v",
			}
		}`, u.Id, u.Username, u.Email)
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
