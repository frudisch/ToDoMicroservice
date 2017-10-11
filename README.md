# ToDo Mircoservice
## General
This is a simple ToDo Microservice for a Blog article (URL upcoming). The service has five endpoints for creating, updating, reading and deleting ToDos from a Postgres database.

## Setup
The project can be simply run by checking out and run ```docker-compose up``` inside the project folder. This is starting the database and the Go webserver.

To run the project without docker-compose. It has to be checkout into the ```GO_PATH``` as usual for Golang projects. To run the microservice, you have to start the database and afterwards the Go server. To start the Postgres database, run the following command: 

```
./db_setup/run.sh
```

To start the Go server, use the following command, which contains environment variables for the database connection:

```
./startServer.sh
```

To run the test cases, you have to navigate to the app subfolder and run inside ```go test -v ```. To successfully run this command, the database needs to be running. The command runs every test case within main_test.go.

## Database
The used database is Postgres inside a Docker container. The Dockerfile uses a Postgres image with an initial SQL file to setup the database and table with the necessary login data. This information could be found in the init.sql file. The two .sh files could be used to create, run and stop the Postgres docker container.

## Example Responses
### GET localhost:8080/todos
Result:
```
[
    {
        "id": 1,
        "name": "GO REST API",
        "description": "Setup Go Rest API for Blog entry",
        "dueTo": 1507273200000
    },
    {
        "id": 2,
        "name": "Create Blog Entry!",
        "description": "Create a awesome Blog entry",
        "dueTo": 1507294800000
    }
]
```

### GET localhost:8080/todo/1
Result:
```
{
    "id": 1,
    "name": "GO REST API",
    "description": "Setup Go Rest API for Blog entry",
    "dueTo": 1507273200000
}
```

### PUT localhost:8080/todo/1
Payload:
```
{
    "name": "GO REST API - update!",
    "description": "Setup Go Rest API for Blog entry, nearly completed",
    "dueTo": 1507273200000
}
```
Result:
```
{
    "id": 1,
    "name": "GO REST API - update!",
    "description": "Setup Go Rest API for Blog entry, nearly complete!",
    "dueTo": 1507273200000
}
```

### POST localhost:8080/todo
Payload:
```
{
    "name": "Test insert",
    "description": "Test insert to delete it afterwards!",
    "dueTo": 1517343200000
}
```
Result:
```
{
    "id": 3,
    "name": "Test insert",
    "description": "Test insert to delete it afterwards!",
    "dueTo": 1617343200000
}
```

### DELETE localhost:8080/todo/3
Result:
```
{
    "result": "success"
}
```