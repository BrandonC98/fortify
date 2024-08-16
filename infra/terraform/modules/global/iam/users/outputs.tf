output "pgp_key" {
  value = {
    public_key  = pgp_key.fortify_user_pgp_key.public_key
    private_key = pgp_key.fortify_user_pgp_key.private_key
  }
  sensitive   = true
  description = "PGP Encryption key used for secure transport of the users secret access key"
}

output "user_key" {
  value = {
    access_key_id     = aws_iam_access_key.fortify_access_key.id
    secret_access_key = aws_iam_access_key.fortify_access_key.encrypted_secret
  }
  sensitive   = true
  description = "Access ID and encrypted secret access key"
}

output "policy" {
  value = {
    arn = aws_iam_policy.fortify_policy.arn
    id  = aws_iam_policy.fortify_policy.id
  }
  description = "Policy ID and ARN"
}

