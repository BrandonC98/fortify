# Generator
This service is apart of the Fortify project. It is a API that generates random text and encrypts the text before returning.

## prerequisites
- Go 
- Docker
- Golangci-linter
- Just

## Commands

- `just run` - run the service
- `just test` - run the tests
- `just build-image` - create a docker image of the service
- `just run-container` - run the docker image
- `just` - list all the just commands with descriptions

## Running the service
This service can be run locally and be run in a docker container

**Locally**
Run the following command to start the service:
```bash
just build
just run-build
```

**Docker**
Run these commands to build the image and run the container:
```bash
just build-image
just run-container
```

### Service endpoints
- `/ping` - healthcheck endpoint, returns a message of pong 
- `/generate` - creates a random text and returns the text encrypted
