# PassMan Pass Gen
This service is apart of the PassMan project. It is a API that generates random password and encrypts them.

## prerequisites
- Go 
- Docker
- Golangci-linter
- Just


## Getting started
This service makes heavy use of a justfile for development.
It's highly recommended to use the just commands over the normal commands as the just commands loads in the `.env` file's environment variables which is needed for some commands to work properly

### Common commands

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
- `/generate` - creates a random password and returns the password encrypted
