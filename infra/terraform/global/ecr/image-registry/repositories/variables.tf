variable "repositories" {
  type        = map(any)
  description = "ECR Repositories. Include name and policy"

  default = {

    fortify_repo = {
      name             = "fortify"
      lifecycle_policy = "policy.json"
    }

    generator_repo = {
      name             = "generator"
      lifecycle_policy = "policy.json"
    }
  }
}
