package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

var u User

func main() {
	lambda.Start(Handler)
}

// Lambda function that creates a new user on DynamoDB with provided username & email
//
// Client must send the "username" and "email" through request body
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Object where we are going to store the sent username and email
	var uCmd UserCMD

	// Object where we are going to store the User
	var u User

	// Save sent values on our Objects
	marshalErr := json.Unmarshal([]byte(req.Body), &uCmd)

	if marshalErr != nil {
		return response("Couldn't unmarshal json into monster struct", http.StatusBadRequest), nil
	}

	// Generate a new uuid
	uuid := uuid.NewString()
	u.Id = uuid
	u.Username = uCmd.Username
	u.Email = uCmd.Email

	// SaveUser controller saves the user in DynamoDB
	dynamoErr := SaveUser(u)

	if dynamoErr != nil {
		return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	}

	// Decode the User object into a json
	decodedUser, _ := json.Marshal(u)

	// Send a response
	json := fmt.Sprintf(`{
		"message": "user created succesfully",
		"user": %v
	}`, string(decodedUser))
	return response(json, http.StatusOK), nil
}

// Send a HTTP response
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
