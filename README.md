# balance
## Microservice for working with user balances
#

## Running the application

```sh
docker-compose up --build
```

### Running the application without Docker

```sh
go run .
```

### Running tests (app must be running)

```sh
go test .
```
#
## HTTP Methods



```
"GET" /ping — Checking the server connection

    example: 
        "GET" :8080/ping
```
#

```
"GET" /balance — Gets the user's balance

    example: 
        "GET" :8080/balance
```
```json
{
    "id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
}
```
#

```
"POST" /user — Create a new user

    example: 
        "POST" :8080/user
```

```json
{
    "funds" : "55.55" // optional
}
```

#
```
"PATCH" /deposit — Depositing funds to the user's balance

    example: 
        "PACTCH" :8080/deposit
```
```json
{
    "id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
    "funds" : "1000"
}
```
#

```
"PATCH" /reserve — Reservation of funds on the user's balance

    example: 
        "PATCH" :8080/reserve
```
```json
{
    "userID" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
    "serviceID" : "ec6741fa-4b02-4e03-a303-0fa96eb15d15",
	"orderID" : "ec6705fa-4b00-0e11-a013-8fa88eb74d35",
	"cost" : 3.5
}
```
#

```
"PATCH" /receipt — Charge off the user's balance

    example: 
        "PATCH" :8080/receipt
```
```json
{
	"userID" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
	"serviceID" : "ec6741fa-4b02-4e03-a303-0fa96eb15d15"
}
```
#

```
"DELETE" /user — Delete user

    example: 
        "DELETE" :8080/
```
```json
{
	"id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
}
```
#