data "aws_ecr_authorization_token" "ecr_token" {
  registry_id = aws_ecr_repository.fortify_ecr.registry_id
}
