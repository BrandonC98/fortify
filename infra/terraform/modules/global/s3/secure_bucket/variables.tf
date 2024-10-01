variable "bucket_name" {
  type        = string
  description = "Name of the s3 bucket"
}

variable "enable_bucket_versioning" {
  type        = string
  description = "Enable bucket versioning(Enable, Suspended or Disable)"
  default     = "Enabled"
}

variable "prevent_bucket_destruction" {
  type        = bool
  description = "Set if the bucket can be destroyed"
  default     = true

}
