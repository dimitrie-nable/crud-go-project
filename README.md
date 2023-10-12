# crud-go-project

This project uses Golang, the Gin framework, MySQL and Docker.

### Running the application
Before running the application a MySQL database must be started, so using the docker-compose,
the database container can be started by running the following command in crud-go-project/build:
```
docker-compose up
```

You can run the application by going in crud-go-project/cmd/crud-go-project
and running the command:
```
go run main.go
```
For testing the application move to crud-go-project/internal/http and run:
```
go test
```
Go will only consider the files that have "_test" at the end of their name for
running the tests. The tests are only for the CRUD endpoints so that is why we have them there.
Also, the tests should be placed in the packet that is going to be tested.

### Prerequisites

Docker\
Go

### Setup local Docker
```
export MYSQL_USER=...
export MYSQL_PASSWORD=...
```

### App functionality
The application mocks cracking passwords of users that consist of only numbers.
It iterates through all the users and tries all the possible numbers for each user, 
until one fits with the one provided. 
Each user gets his own goroutine so the iterations for guessing the password should be done
concurrently.

Accessing each functionality of the app: create user, get users, get passwords etc., 
happens using an endpoint.