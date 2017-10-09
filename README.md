# ToDo Mircoservice
## General
This is a simple ToDo Microservice for a Blog article (URL upcoming). The service has five endpoints for creating, updating, reading and deleting ToDos from a Postgres database.

## Setup
The project has to be checkout into the ```GO_PATH``` as usual for Golang projects. To Run the microservice, you have to start the database and afterwards the Go server. To start the Postgres database, docker needs to be running and you can use the following command: 

```
./db_setup/run.sh
```

To start the Go server, use the following command:

```
go run main.go
```

To run the testcases, you have to navigate to the app subfolder and run inside ```go test -v ```. This command runs every test case within main_test.go

## Database
The used database is Postgres inside a Docker container. The Dockerfile uses a Postgres image with an initial SQL file to setup the database and table with the necessary login data. This information could be found in the init.sql file. The two .sh files could be used to create, run and stop the Postgres docker container. 