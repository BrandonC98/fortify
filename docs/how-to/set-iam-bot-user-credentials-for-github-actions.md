# Set IAM bot user credentials for GitHub actions 

1. Go to the`infra/scripts` directory.

2. Run the `set-gitub-secrets-for-ecr.sh` script
```bash
./set-gitub-secrets-for-ecr.sh
```

3. Check the Secrets have been updated. This can be done by listing the secrets and when it was last updated. Note the value can't be printed through the `gh` and it should not be attempted to be printed in a action.
```bash
gh secret list --json name --json updatedAt | jq '.[] | select( .name == "ECR_USER_ACCESS_KEY_ID" or  .name == "ECR_USER_SECRET_ACCESS_KEY" )'
```

## Explanation
The `set-github-secrets-for-ecr.sh` script will access the Terraform ouput for the IAC code the user and use it to get multiple values:
- GPG Private Key - The Private key needs to be imported to the keyring so we can in a future step decrypt the Secret access key. This needs to be done as the Terraform handles generating the GPG key.
- Access Key ID - This taken from the Terraform output and set as the GitHub secret
- Secret Access Key - This value is outputed by Terraform encrypted and base64 encoded. It is first decoded from base64, then the GPG private key is used to decrypt it. The decrypted value then set as the GitHub secret value

## When and Why
These steps should only need to be executed if the bot user has been changed to a different/new user.

## Good to know
- This script is idempotent. If Secrets are already set they will be updated.
