resource "aws_iam_policy" "fortify_policy" {
  name        = var.policy_name
  description = var.policy_description

  policy = fileexists(var.policy_path) ? file(var.policy_path) : file("${path.module}/${var.policy_path}")

  tags = {
    name    = var.policy_name
    project = "fortify"
  }
}
