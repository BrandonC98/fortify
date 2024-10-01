# Create New Image Repository 

1. Go to the`infra/terraform/global/ecr/image-registry/repositories` directory.

2. Open the `variables.tf` file

3. A new entry needs to be made in the default map for `repositories` variable
    3.1. The `<NAME>_repo` is how terraform will refer to this object. This needs to be unique from any other key in the default map.
    3.2. The `<SERVICE_NAME>` is used as the ecr repository name in aws. This will be shown in the image reppositories url
    3.3. `<POLICY_FILE_PATH>` is used to attach lifecycle polcies to the repository. The path needs to be relative to the module folder or the `repositories` folder. Use `policy.json` for most situaations.
```tf
<NAME>_repo = {
    name = "<SERVICE_NAME>"
    lifecycle_policy = "<POLICY_FILE_PATH>"
}
```

4. If a new policy has been created in the module directory update terraform
```
terraform get -update
```

5. Run terraform apply to create the repository
```bash
terraform apply
```

## When and Why
Use these steps when a new service needs to store it's docker image

## Good to know
- The policy can be set to `null` if none is needed
- Get the terraform output for a single repository `terraform output -json | jq.repositories_data.value.<NAME>_repo` 
