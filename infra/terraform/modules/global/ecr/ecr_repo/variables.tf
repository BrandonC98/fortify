variable "repo_name" {
  type        = string
  description = "ECR Repository name"
}

variable "lifecycle_policy_file" {
  type        = string
  description = "path of the lifecycle policy file to apply to ECR"
  default     = null
}
