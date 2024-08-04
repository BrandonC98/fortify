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

resource "aws_ecr_registry_scanning_configuration" "fortify_ecr_scanning" {
  scan_type = var.scan_type
}
