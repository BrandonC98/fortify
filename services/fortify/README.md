# Fortify
This service provides the web interface for the Fortify secret manager.

## Rerequisites
- Go
- Docker
- Golangci-linter
- Just
- mysql

## Getting Started 
This service makes heavy use of a justfile for development. It's highly recommended to use the just commands over the normal commands as the just commands loads in the .env file's environment variables which is needed for some commands to work properly
### Common Commands
- `just run` - run the service
- `just test` - run the tests
- `just build-image` - create a docker image of the service
- `just run-container` - run the docker image
- `just` - list all the just commands with descriptions


## Running The Service
It's recommended to use docker-compose to run the service locally as it will provide the database.
```bash
just build-image
just run-container
```
