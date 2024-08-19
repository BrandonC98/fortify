provider "aws" {
  region = "eu-west-2"
}

terraform {
  backend "s3" {
    bucket = "fortify-backend"
    key    = "global/s3/backend-bucket/terraform.tfstate"
    region = "eu-west-2"

    dynamodb_table = "fortify-tf-lock"
    encrypt        = true

  }
}

module "backend_bucket" {
  source      = "git@github.com:BrandonC98/fortify.git//infra/terraform/modules/global/s3/secure_bucket"
  bucket_name = "fortify-backend"
}

# DynamoDB for locking using key-value store
resource "aws_dynamodb_table" "tf_lock" {
  name         = "fortify-tf-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
