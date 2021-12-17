How to test.
1. DYNAMODB: Create table called rcs-serverless-users.
2. LAMBDA: Upload each microservice on a separate function. Remember to change "hello" for "main" as function package and to select a role with proper policies in order to provide access to dynamodb to the function.
3. API GATEWAY: Create an API REST. Create 2 resources: users & user. Remember to provide proxy and to test.