terraform {
  required_providers {
    pgp = {
      version = "0.2.4"
      source  = "ekristen/pgp"
    }
  }
}

locals {
  policy_name = trimsuffix(basename(var.policy_path), ".json")
}

resource "aws_iam_user" "fortify_user" {
  name = var.user_name

  tags = {
    name    = var.user_name
    project = "fortify"
  }
}

resource "aws_iam_access_key" "fortify_access_key" {
  user    = aws_iam_user.fortify_user.name
  pgp_key = pgp_key.fortify_user_pgp_key.public_key_base64
}

resource "aws_iam_user_policy" "fortify_user_policy" {
  name   = local.policy_name
  user   = aws_iam_user.fortify_user.name
  policy = aws_iam_policy.fortify_policy.policy
}

# Pass in a json policy file to attach to the user
resource "aws_iam_policy" "fortify_policy" {
  name        = local.policy_name
  description = var.policy_description

  policy = fileexists(var.policy_path) ? file(var.policy_path) : file("${path.module}/${var.policy_path}")

  tags = {
    name    = local.policy_name
    project = "fortify"
  }
}

# Generate a pgp key for the user to encrypt their secrets
resource "pgp_key" "fortify_user_pgp_key" {
  name    = var.user_name
  email   = var.user_email
  comment = "PGP key for ${var.user_name}"
}
