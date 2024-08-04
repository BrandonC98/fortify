resource "aws_ecr_lifecycle_policy" "fortify_ecr_lifecycle_policy" {
  repository = var.repo_name
  policy     = file("${path.module}/${var.lifecycle_policy_file}")
}
