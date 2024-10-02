output "repositories_data" {
  value = {
    for k, repo in module.image_registry : k => {
      repository = k
      name       = var.repositories[k].name
      url        = repo.repo_url
    }
  }
}

output "session_tokens" {
  value = {
    for k, repo in module.image_registry : k => {
      token = repo.session_token
    }
  }

  description = "Session tokens for ecr repositories"
  sensitive   = true
}
