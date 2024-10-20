# Set Kubernetes Secrets 

1. Open the `infra/kubernetes/deployments/secrets/` directory
2. run the `create-secret-env-file.sh` script. This will create a `secret-env` file with empty secret fields
3. Add secrets to the `secret-env` file
4. Create the secret object by running either of the following commands
    1. Deploy just the secrets: `kubectl create secret generic fortify-secrets --from-env-file=secret-env`
    1. Deploy all resources, run the following command in the parent directory: `bash deploy-objects-to-cluster.sh`

## When and Why
- This is required for running the service in kubernetes, resources will crashloop if secrets aren't set

## Good to know
- `secret-env` is ignored by git so won't be commited to github
- If the `encryption_key` is changes while the applications are running then values encrypted using the old key cannot be decrypted
