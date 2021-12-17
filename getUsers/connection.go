package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Table from DynamoDB
const tableName = "rcs-serverless-users"

// Controller dedicated to fetch the users.
func FetchUsers() (arr []User, err error) {
	// Stablish the connection with DynamoDB
	dynamoSession := createDynamoSession()

	ctx := context.Background()

	// Get the users from the database
	scanOutput, err := dynamoSession.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		fmt.Println("Failed to get from dynamo: ", err.Error())
		return
	}

	for _, item := range scanOutput.Items {
		var u User
		err := dynamodbattribute.UnmarshalMap(item, &u)
		if err != nil {
			fmt.Println("Failed to unmarshalmap from dynamo: ", err.Error())
			return nil, err
		}
		arr = append(arr, u)
	}

	return
}

// Stablish the connection to DynamoDB
func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}
