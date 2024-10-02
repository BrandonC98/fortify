# Push Images to ECR

1. Make sure `gh` is setup
2. Run the following command to start the Github action that will push a Docker image to ECR. 
    2.1 Replace `<SERVICE_NAME>` with the name of the service to push to ECR.
    2.2 Replace `<TAG>` with the image tag.
    2.3 Replace `<ENV_FILE>` with the env file name to get the configuration values from. These files can be found in the `configuration` folder within the folder of the service you are targetting.

```bash
gh run <SERVICE_NAME>-push-image.yml -F image-tag=<TAG> -F environment-file=<ENV_FILE>
```

## When and Why
Used when new version of a service's docker image should be stored in ECR for use in running the applications

## Good to know
Currently supported serivces
- `generator`
- `fortify`

This can also be triggered by running Action in the GitHub repository
