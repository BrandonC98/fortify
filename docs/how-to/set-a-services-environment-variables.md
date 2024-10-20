# Set a service's environment variables 
The steps for setting up a service's environment variables will depend on where you plan to the build

**Local(no docker)**
When running locally without docker then environment variables can be put in the `.env` file for the service. Then run the service using the `just` command, which will load in `.env` file and set the environment variables in the file.

**Local(docker)**
For docker and docker-compose builds then environment variables use `.env` files inside the `services/<SERVICE_NAME>/configuration/` directory. These files can be passed into the build command or set in the `compose.yaml` file.

**GitHub Action**
The build action can has an paramater for passing in `.env` file names. This will search the `services/<SERVICE_NAME>/configuration/` directory for the file. By default this job will have the value set to `None` which is an empty file, so no envrionment variables are used.

`None` is used for kubernetes builds so `configmaps` can be used to set the envrionment variables

## Good to know
- For Docker related scenarios the `.env` passed in will be applied at the last step of the dockerfile during the `setup.sh` script
