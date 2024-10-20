# Run k8s Cluster 

Needed tools:
- kops
- kubectl
- aws credentials
- just

## Cluster Setup
- Cluster backends are created in a S3 bucket for the project. The following command can be used to create the cluster's backend
```bash
just create-cluster
```
- The Cluster can then be deployed with the following command. This will also run validation to make sure the cluster has been setup properly, this will take 10 minutes
```bash
just update-cluster
```

- The Cluster can be deleted with this command. This will delete the deployment and remove the backend from s3
```bash
just delete-cluster
```

- Replace the Cluster configuration with the local configuration, using this command
```bash
just replace-cluster
```
## Deploying into the Cluster
1. Once the Cluster is fully setup then resources can be deployed into the environment.
2. Go to the `infra/kubernetes/deployments`
3. Make sure secrets have been set. check for [set-k8s-secrets.md](./set-k8s-secrets.md) on how to do this
4. Run `deploy-objects-to-cluster.sh`. This will deploy all the resources into the cluster

## When and Why
- Each just command will required confirmation before running for safety. confirm by typing `y` or `yes`
- Run the create and update command to raise the cluster
- Run the delete command to remove the cluster once finished with it
- Regular kops commands will work but the kops commands are best to be used for common use cases to prevent errors

## Good to know
- Clusters are created with kops
- Each cluster will have its own file in the `infra/kubernetes/kops-clusters/` directory
    - Currently the only cluster is `prod`
- Clusters are defined using configuration yaml files, if the cluster needs to be updated the yaml files are where the change should take place
- Common Configuration files:
    - `cluster.yaml` - Contains setup around the cluster whole cluster and tools
    - `master-instance-group.yaml` - Defines configuration around the master/control plane instance
    - `nodes-instance-group.yaml` - Defines configuration around the each node instance
- Access web application in a browser by using the `fortify-ingress` address
