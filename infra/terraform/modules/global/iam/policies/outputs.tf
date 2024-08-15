output "policy_arn" {
  value = aws_iam_policy.fortify_policy.arn
}

output "policy_id" {
  value = aws_iam_policy.fortify_policy.id
}

output "policy_attachment_count" {
  value = aws_iam_policy.fortify_policy.attachment_count
}
