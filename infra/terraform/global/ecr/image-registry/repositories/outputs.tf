output "repositories_data" {
  value = {
    for k, repo in module.image_registry : k => {
      repository = k
      name       = var.repositories[k].name
      url        = repo.repo_url
    }
  }
}
