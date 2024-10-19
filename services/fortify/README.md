# Fortify
This service provides the web interface for storing key/value pair.

## requisite
- Go
- Docker
- Golangci-linter
- Just
- mysql

## Getting Started 
It's highly recommended to use the just commands over the normal commands as the just commands loads in the .env file's environment variables which is needed for some commands to work properly
### Commands
- `just run` - run the service
- `just test` - run the tests
- `just build-image` - create a docker image of the service
- `just run-container` - run the docker image
- `just` - list all the just commands with descriptions


## Running The Service
It's recommended to use docker-compose to run the service locally as it will provide the database which is required for running the serivce.
```bash
just build-image
just run-container
docker-compose up
```
