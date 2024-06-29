# Docker
This project uses Docker to containerize each application.

Each docker file uses builder images to build the source code, then the executable is put on a seperate image. This is so the container has the smallest possible storage size. Alpine is also used as a base image because of it's slim size.

This project also uses docker compose to run multiple images together. healthchecks are run for applications that depend on others. They also make great use of .env files to customize the environment varaible set on the containers.
