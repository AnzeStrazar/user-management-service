# user-management-service

This API contains a simple implementation of user and group management service.

### To get everything running, follow these steps:

1. By running `docker-compose up` we pull & build all required images and bring up all the
containers for the application. API is exposed on port: 8000.

2. In order to use the application, you'll need to provide the state and structure to the database.
To do this you need to run: `go run . --migrate` in the project directory, this will run all database migrations.
To clear the database, you need to shred the state. In order to do that, you need to run: `go run . --migrate-down`.

### To get everything running localy, follow these steps:
_Note: You will need latest release of Go installed on your machine in order to run API server locally._

1. By running `docker-compose up -d` we pull & build all required images and bring up all the
containers for the application. API is exposed on port: 8080.

2. In order to use the application, you'll need to provide the state and structure to the database.
To do this you need to run: `go run . --migrate` in the project directory, this will run all database migrations.
To clear the database, you need to shred the state. In order to do that, you need to run: `go run . --migrate-down`.

3. By running `go run .` we run HTTP API Server.

### OpenAPI docs
API documentation is described in `docs/openapi.yaml`.  
This file is used when you run `docker-compose up -d` and go to `http://localhost:5050/docs`.

### Testing
If not already, you first need to execute `docker-compose up -d` and then `go run . --migrate`.

To run tests you need to be in `tests` directory.
You can run tests using `go test -v`.

Or simply run `go test ./...` in the project directory.

### Configuration

In `config.json` we define all necessary information, needed for connection to Postgres database.