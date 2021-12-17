# AWS - LAMBDA - TESTS
I made this project in order to test how AWS Lambda, DynamoDB and API Gateway work and how may i work with them through Go language.
What i made is a simple CRUD of users where you can [ create, fetch by id, fetch all and delete by id ] users.


## How to run the project:


### 1. DYNAMODB
Create table.

- rcs-serverless-users.

### 2. LAMBDA
Upload each microservice on a separate function. Remember to change "hello" for "main" as function package and to select a role with proper policies in order to provide access to dynamodb to the function.

- rcs-users-fetchall
- rcs-user-fetchbyid
- rcs-user-deletebyid
- rcs-user-create
    
### 3. API GATEWAY
Create an API REST. Create 2 resources: users & user. Remember to provide proxy and to test.

- rcs-users
    
    
    /user
    - [DELETE] -> rcs-user-deletebyid
        Must send id from the record you want to delete as queryparam.
                
    - [GET] -> rcs-user-fetchbyid
        Must send id from the record you want to fetch as queryparam.
                
    - [POST] -> rcs-user-create
        Must send "username" & "email" through the body on a JSON object.
            
    /users
    - [GET] -> rcs-users-fetchall
