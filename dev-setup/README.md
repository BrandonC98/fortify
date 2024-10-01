# Dev Setup
Outlines the needed tools to develop and run the Fortify project.

## Tools
- go version 1.22
- golangci-linter
- just
- docker

## Standard Command Layer
This project uses just as a standard command layer. This means just commands can be used throughout the repo for common tasks. e.g. `just run` can be used in a Go project to run the application, the same command could be used for a python application. This means as long as the user has all the prerequites installed they won't need to familiar with the specific projects details to achive basic functionality

Just will also handle injecting environment variables from `.env` files. Using the raw command instead of the just alternative may cause errors if the command requires envirnoment varaibles
