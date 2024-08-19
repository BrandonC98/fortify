provider "aws" {
  region = "eu-west-2"
}

terraform {
  backend "s3" {
    bucket = "fortify-backend"
    key    = "global/iam/users/ecr-user/terraform.tfstate"
    region = "eu-west-2"

    dynamodb_table = "fortify-tf-lock"
    encrypt        = true
  }
}

module "ecr_user" {
  source = "git@github.com:BrandonC98/fortify.git//infra/terraform/modules/global/iam/users"

  user_name  = var.user_name
  user_email = var.user_email

  policy_path        = var.policy_path
  policy_description = var.policy_description
}
