#!/bin/bash

cd ../terraform/global/iam/users/ecr-user

echo "Importing ecr-user's gpg private key"
terraform output -json pgp_key | jq -r .private_key | gpg --import

echo "Setting AWS_ACCESS_KEY_ID secret for GitHub actions"
terraform output -json user_key | jq -r .access_key_id | gh secret set AWS_ACCESS_KEY_ID -a actions -b -

echo "Setting AWS_SECRET_ACCESS_KEY secret for GitHub actions"
terraform output -json user_key | jq -r .secret_access_key | base64 --decode | gpg --decrypt -q | gh secret set AWS_SECRET_ACCESS_KEY -a actions -b -


