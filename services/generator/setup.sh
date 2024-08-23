#!/bin/sh
# First element is the file name of the env file which will be used to set the configuration for the container
# All other arguments passed into the script will be used to start the application

ENV_FILE=$1

# Shift the arg list one place to the right. This removes the first element from the list so when the list is executed it won't include the env file
shift 

echo "Setting environment variables"
if [ -f "configuration/$ENV_FILE" ]; then

	# loop through env file exporting any line that isn't a comment
	export $(grep -v '^#' configuration/$ENV_FILE | xargs)

else
	echo "configuration/$ENV_FILE not found"
fi

echo "Starting application with command: $@"
exec $@
