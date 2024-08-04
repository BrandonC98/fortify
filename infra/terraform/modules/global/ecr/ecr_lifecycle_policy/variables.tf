variable "repo_name" {
  type        = string
  description = "Name of the ECR repository"
}

variable "lifecycle_policy_file" {
  type        = string
  description = "Name of the lifecycle policy file to apply to ECR"
  default     = "policy.json"
}
