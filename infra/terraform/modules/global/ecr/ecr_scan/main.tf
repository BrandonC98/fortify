resource "aws_ecr_registry_scanning_configuration" "fortify_ecr_scanning" {
  scan_type = var.scan_type

  rule {
    scan_frequency = "SCAN_ON_PUSH"
    repository_filter {
      filter      = var.scan_filter
      filter_type = "WILDCARD"
    }

  }
}
