resource "aws_ecr_repository" "fortify_ecr" {
  name = var.repo_name

  encryption_configuration {
    encryption_type = "AES256"
  }

  tags = {
    name    = var.repo_name
    project = "fortify"
  }
}

resource "aws_ecr_lifecycle_policy" "fortify_ecr_lifecycle_policy" {
  count      = var.lifecycle_policy_file != null ? 1 : 0
  repository = aws_ecr_repository.fortify_ecr.name
  policy     = fileexists(var.lifecycle_policy_file) ? file(var.lifecycle_policy_file) : file("${path.module}/${var.lifecycle_policy_file}")
}
