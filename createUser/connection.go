package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB Table name
const tableName = "rcs-serverless-users"

// Controller dedicated to save the user, should receive the User as a parameter.
func SaveUser(user User) error {
	// Store the user values as dynamodb.AttributeValue but as
	// dynamodbattribute.MarshalMap(u) was returning an empty map
	// We create it manually
	av := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(user.Id),
		},
		"username": {
			S: aws.String(user.Username),
		},

		"email": {
			S: aws.String(user.Email),
		},
	}

	// Stablish the connection with DynamoDB
	dynamoSession := createDynamoSession()

	ctx := context.Background()

	// Insert the User on the database
	_, writeErr := dynamoSession.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})

	if writeErr != nil {
		fmt.Println("Failed to write to dynamo: ", writeErr.Error())
		return writeErr
	}

	return nil
}

// Stablish the connection with DynamoDB
func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}
