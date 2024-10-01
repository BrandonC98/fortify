provider "aws" {
  region = "eu-west-2"
}

terraform {
  backend "s3" {
    bucket = "fortify-backend"
    key    = "global/ecr/image-registry/repositories/terraform.tfstate"
    region = "eu-west-2"

    dynamodb_table = "fortify-tf-lock"
    encrypt        = true
  }
}

module "image_registry" {
  source = "git@github.com:BrandonC98/fortify.git//infra/terraform/modules/global/ecr/ecr_repo"

  for_each = var.repositories

  repo_name             = each.value.name
  lifecycle_policy_file = each.value.lifecycle_policy
}
