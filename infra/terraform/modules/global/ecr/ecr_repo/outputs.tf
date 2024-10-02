output "registry_id" {
  value = aws_ecr_repository.fortify_ecr.registry_id
}
output "repo_url" {
  value = aws_ecr_repository.fortify_ecr.repository_url
}

output "session_token" {
  value = data.aws_ecr_authorization_token.ecr_token
}
