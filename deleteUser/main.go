package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Lambda function that deletes a user from DynamoDB with provided ID
//
// Client must send the "id" through the urlparams
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse,
	error) {

	// Get id from url parameters
	id := request.QueryStringParameters["id"]

	// Delete user is the controller dedicated to delete the user
	u, err := DeleteUser(id)
	if err != nil {
		json := fmt.Sprintf(`{
			"message": "there was an error at the server",
			"provided_id": "%v"
		}`, id)
		return response(json, http.StatusInternalServerError), nil
	}

	// If there are no errors but it could not find any user to delete.
	if u.Username == "" {
		json := fmt.Sprintf(`{
			"message": "could not fetch any user with provided id",
			"provided_id": "%v"
		}`, id)
		return response(json, 200), nil
	}

	userJson, _ := json.Marshal(u)

	// If we find a user, send it through the response body.
	json := fmt.Sprintf(`{
			"message": "user deleted succesfully",
			"user": %v
		}`, string(userJson))
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
