variable "repo_name" {
  type        = string
  description = "ECR Repository name"
}

variable "scan_type" {
  type        = string
  description = "Type of d scan to run"
  default     = "BASIC"
}
