
provider "aws" {
  region = "eu-west-2"
}

terraform {
  backend "s3" {
    bucket = "fortify-backend"
    key    = "global/ecr/image-registry/security-scanning/terraform.tfstate"
    region = "eu-west-2"

    dynamodb_table = "fortify-tf-lock"
    encrypt        = true

  }
}

module "fortify_ecr_scanning" {
  source = "git@github.com:BrandonC98/fortify.git//infra/terraform/modules/global/ecr/ecr_scan"
}
