output "pgp_key" {
  value       = module.ecr_user.pgp_key
  sensitive   = true
  description = "PGP Encryption key used for secure transport of the users secret access key"
}

output "user_key" {
  value       = module.ecr_user.user_key
  sensitive   = true
  description = "Access ID and encrypted secret access key"
}

output "policy" {
  value       = module.ecr_user.policy
  description = "Policy ID and ARN"
}

