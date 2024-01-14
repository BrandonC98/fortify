set dotenv-load 

default:
	just --list

# Run test suite in verbose mode with the gin server mode set to test
test $GIN_MODE="test":
	go test ./... -v

run:
	go run ./...

# Build the bin and place it at the relative path of bin/app
build:
	go build -v -o bin/app ./...

# Run the build. pass in a paramater to set the gin server mode. default is debug
run-build $GIN_MODE="debug":
	./bin/app

# Run linter
lint:
	golangci-lint run ./... --verbose

# Generate a .env file
gen-dotenv mode="debug" port="8080" min="7" max="25":
	touch .env
	echo "GIN_MODE=\"{{mode}}\"" >> .env
	echo "PASSMAN_PORT=\"{{port}}\"" >> .env
