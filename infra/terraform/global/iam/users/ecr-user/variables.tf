variable "user_name" {
  description = "Name of the user"
  default     = "ecr-bot-user"
}

variable "user_email" {
  description = "Email address of the user"
  default     = "brandoncampbell98@hotmail.co.uk"
}

variable "policy_path" {
  description = "Path to the policy file"
  default     = "ecr-policy.json"
}

variable "policy_description" {
  description = "Description of the policy"
  default     = "Allows the user to push and pull images to the ECR service"
}
